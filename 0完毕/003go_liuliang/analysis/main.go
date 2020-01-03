package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mediocregopher/radix.v2/pool"  // redis pool池
	"github.com/mediocregopher/radix.v2/redis" //redis操作
	"github.com/mgutz/str"                     //字符串格式化
	"github.com/sirupsen/logrus"               //打日志
)

const handleDig = " /dig?"
const handleMovie = "/movie/"
const handleList = "/list/"
const handleHTML = ".html"

var log = logrus.New()
var redisCli redis.Client

type cmdParams struct {
	logFilePath string
	routineNum  int
}
type digData struct {
	time  string
	url   string
	refer string
	ua    string
}
type urlData struct {
	data  digData
	uid   string
	unode urlNode
}
type urlNode struct {
	unType string //详情页 列表页 首页
	unRid  int    // Resource资源id
	unURL  string // 当前这个页面的url
	unTime string // 当前访问这个页面的时间
}
type storageBlock struct {
	counterType  string
	storageModel string
	unode        urlNode
}

func readFileLineByLine(Params cmdParams, logChannel chan string) (err error) {

	fd, err := os.Open(Params.logFilePath)
	defer fd.Close()
	if err != nil {
		log.Warningf("readFileLineByLine os.Open(Params.logFilePath) open %s error", Params.logFilePath)
		return err
	}
	count := 0
	bufferRead := bufio.NewReader(fd)
	for {
		line, err := bufferRead.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				time.Sleep(3 * time.Second)
				log.Infof("readFileLineByLine wait %d second", 3)
			} else {
				log.Warningf("readFileLineByLine err = %s", err)
			}
			return err
		}
		logChannel <- line
		// log.Infof("%s", line)
		count++
		if count%(1000*Params.routineNum) == 0 {
			log.Infof("readFileLineByLine line %d", count)
		}
		return err
	}

}

func logConsumer(logChannel chan string, pvChannel, uvChannel chan urlData) error {
	for logStr := range logChannel {
		// 切割日志字符串，扣出打点上报的数据
		data := cutLogFetchData(logStr)
		//uid
		// 说明 课程模拟生成uid md5(refer+ua)
		hasher := md5.New()
		hasher.Write([]byte(data.refer + data.ua))
		uid := hex.EncodeToString(hasher.Sum(nil))

		// 很多解析的工作都有可以放在这里完成
		urlNode := formatURL(data.url, data.time)
		uData := urlData{
			data:  data,
			uid:   uid,
			unode: urlNode,
		}
		// log.Infoln(uData)
		pvChannel <- uData
		uvChannel <- uData
	}
	return nil
}

func cutLogFetchData(logStr string) digData {
	logStr = strings.TrimSpace(logStr)

	pos1 := str.IndexOf(logStr, handleDig, 0)
	if pos1 == -1 {
		return digData{}
	}
	pos1 += len(handleDig)
	pos2 := str.IndexOf(logStr, "200 43", pos1)
	d := str.Substr(logStr, pos1, pos2-pos1)
	urlInfo, err := url.Parse("http:127.0.0.1?" + d)
	if err != nil {
		return digData{}
	}
	data := urlInfo.Query()
	return digData{
		time:  data.Get("time"),
		url:   data.Get("url"),
		refer: data.Get("refer"),
		ua:    data.Get("ua"),
	}
}
func formatURL(url, t string) urlNode {
	pos1 := str.IndexOf(url, handleMovie, 0)
	if pos1 != -1 {
		pos1 += len(handleMovie)
		pos2 := str.IndexOf(url, handleHTML, 0)
		idStr := str.Substr(url, pos1, pos2-pos1)
		id, _ := strconv.Atoi(idStr)
		return urlNode{
			unType: "movie",
			unRid:  id,
			unURL:  url,
			unTime: t,
		}
	} else {
		pos1 := str.IndexOf(url, handleList, 0)
		if pos1 != -1 {
			pos1 += len(handleList)
			pos2 := str.IndexOf(url, handleHTML, 0)
			idStr := str.Substr(url, pos1, pos2-pos1)
			id, _ := strconv.Atoi(idStr)
			return urlNode{
				unType: "list",
				unRid:  id,
				unURL:  url,
				unTime: t,
			}
		} else {
			return urlNode{
				unType: "home",
				unRid:  1,
				unURL:  url,
				unTime: t,
			}
		} // 如果页面很多种，就在这里不断扩展
	}

}

func pvCounter(pvChannel chan urlData, storageChannel chan storageBlock) {
	for data := range pvChannel {
		sItem := storageBlock{
			counterType:  "pv",
			storageModel: "ZINCRBY",
			unode:        data.unode,
		}
		storageChannel <- sItem
	}
}

func uvCounter(uvChannel chan urlData, storageChannel chan storageBlock, redisPool *pool.Pool) {
	for data := range uvChannel {
		// HyperLoglog redis
		HyperLoglogKey := "uv_hpll_" + getTime(data.data.time, "day")
		ret, err := redisPool.Cmd("PFADD", HyperLoglogKey, data.uid, "EX", 86400).Int()
		if err != nil {
			log.Warningf("uvCount redis HyperLoglogKey err = %s", err)
		}
		if ret != 1 {
			continue
		}
		sItem := storageBlock{
			counterType:  "uv",
			storageModel: "ZINCRBY",
			unode:        data.unode,
		}
		storageChannel <- sItem
	}
}

func dataStorage(storageChannel chan storageBlock, redisPool *pool.Pool) {
	for block := range storageChannel {
		preFix := block.counterType + "_"
		// 逐层添加，加洋葱皮的过程
		// 维度：天-小时-分钟
		// 层级：顶级-大分类-小分类-终极页面
		// 存储模型 Redis SortedSet
		setKeys := []string{
			preFix + "day_" + getTime(block.unode.unTime, "day"),
			preFix + "hour_" + getTime(block.unode.unTime, "hour"),
			preFix + "min_" + getTime(block.unode.unTime, "min"),
			preFix + block.unode.unType + "_day_" + getTime(block.unode.unTime, "day"),
			preFix + block.unode.unType + "_hour_" + getTime(block.unode.unTime, "hour"),
			preFix + block.unode.unType + "_min_" + getTime(block.unode.unTime, "min"),
		}
		rowID := block.unode.unRid
		for _, key := range setKeys {
			ret, err := redisPool.Cmd(block.storageModel, key, 1, rowID).Int()
			if ret <= 0 || err != nil {
				log.Errorln("dataStorage redis Storage error ", block.storageModel, key, rowID)
			}
		}
	}
}

func getTime(logTime, timeType string) string {
	var item string
	switch timeType {
	case "day":
		item = "2006-01-02"
	case "hour":
		item = "2006-01-02 15"
	case "min":
		item = "2006-01-02 15:04"
	}
	t, _ := time.Parse(item, time.Now().Format(item))
	return strconv.FormatInt(t.Unix(), 10)
}

func init() {
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)
}

func main() {
	// 接受参数
	logFilePath := flag.String("logFilePath", "../dig.log", "log")
	routineNum := flag.Int("routineNum", 5, "routineNum by goroutine")
	l := flag.String("l", "../tmp/log", "this log")
	flag.Parse()

	params := cmdParams{logFilePath: *logFilePath, routineNum: *routineNum}
	// 打日志
	logFd, err := os.OpenFile(*l, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer logFd.Close()

	if err == nil {
		log.Out = logFd
	} else {
		fmt.Println("err = ", err)
	}
	log.Infoln("Exec start.")
	log.Infof("params:logFilePath=%s,routineNum=%d", params.logFilePath, params.routineNum)

	// 初始化一些channel，用于数据传递
	logChannel := make(chan string, params.routineNum*3)
	pvChannel := make(chan urlData, params.routineNum)
	uvChannel := make(chan urlData, params.routineNum)
	storageChannel := make(chan storageBlock, params.routineNum)

	// redis Pool
	redisPool, err := pool.New("tcp", "127.0.0.1:6379", params.routineNum*2)
	if err != nil {
		log.Fatalln("redis pool created failed")
		panic(err)
	} else {
		go func() {
			for {
				redisPool.Cmd("PING")
				time.Sleep(3 * time.Second)
			}
		}()
	}

	// 日志消费者
	readFileLineByLine(params, logChannel)
	// 创建一组日志处理
	for i := 0; i < params.routineNum; i++ {
		go logConsumer(logChannel, pvChannel, uvChannel)
	}
	// 创建PV UV 统计器
	go pvCounter(pvChannel, storageChannel)
	go uvCounter(uvChannel, storageChannel, redisPool)
	// 创建 存储器
	go dataStorage(storageChannel, redisPool)
	time.Sleep(1000 * time.Second)
}
