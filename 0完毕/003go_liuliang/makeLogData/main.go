package main

import (
	"fmt"
	"flag"
	// "io/ioutil"
	"os"
	"strings"
	"strconv"
	"net/url"
	"math/rand"
	"time"
)

type resource struct{
	url string
	target string
	start int
	end int
}
var uaList  = []string{
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.80 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.163 Safari/535.1",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:6.0) Gecko/20100101 Firefox/6.0",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Opera/9.80 (Windows NT 6.1; U; zh-cn) Presto/2.9.168 Version/11.50",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; InfoPath.3)",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; ) AppleWebKit/534.12 (KHTML, like Gecko) Maxthon/3.0 Safari/534.12",
}

func ruleResource() []resource  {
	var res []resource
	rs1 := resource{
		url:"http://www.hax.com/",
		target:"",
		start:0,
		end:0,
	}
	rs2 := resource{
		url:"http://www.hax.com/index/product/index/cid/7/id/{$id}.html",
		target:"{$id}",
		start:24,
		end:29,
	}
	rs3 := resource{
		url:"http://www.hax.com/index/product/detail/cid/7/id/24/aid/{$id}.html",
		target:"{$id}",
		start:8,
		end:42,
	}
	res = append(res,rs1,rs2,rs3)
	return res
}

func buildUrl(res []resource) []string {
	var list []string
	for _, val := range res {
		if val.target == "" {
			list = append(list,val.url)
		}else{
			for j := val.start; j <= val.end; j++ {
				urlStr := strings.Replace(val.url,val.target,strconv.Itoa(j),-1)
				list = append(list,urlStr)
			}
		}
	}
	return list
}

func makeLog(currentUrl, referUrl, ua string) string {
	u := url.Values{}
	u.Set("time", "1")
	u.Set("url", currentUrl)
	u.Set("refer", referUrl)
	u.Set("ua", ua)
	paramStr := u.Encode()

	logTemplate := "192.168.233.1 - - [02/Oct/2019:05:10:30 -0700] \"GET /dig?{$paramStr}\" 200 43 \"{$Url}\" \"{$ua}\" \"-\" \n"
	log := strings.Replace(logTemplate, "{$paramStr}",paramStr,-1)
	log = strings.Replace(log, "{$Url}",currentUrl,-1)
	log = strings.Replace(log, "{$ua}",ua,-1)
	// fmt.Println(log)
	return log
}

func randInt(min,max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if min > max {
		return max
	}
	return r.Intn(max-min) + min
}

func main()  {
	total := flag.Int("total", 100, "how many rows")
	filePath := flag.String("filePath","E://009-Go/goProjects/src/go_liuliang/run/dig.log","dayin rizhi")
	flag.Parse()
	fmt.Println(*total, *filePath)
	// 需要构造出真实的网站url集合
	res := ruleResource()
	var list []string
	list = buildUrl(res)
	// fmt.Println(list)
	// 按照要求，生成$total 行日志内容，源自url集合
	logStr := ""
	for i := 0; i < *total; i++ {
		currentUrl := list[randInt(0,len(list)-1)]
		referUrl := list[randInt(0,len(list)-1)]
		ua := uaList[randInt(0,len(uaList)-1)]
		logStr += makeLog(currentUrl, referUrl, ua)
		// ioutil.WriteFile(*filePath, []byte(logStr),0644)
	}
	f,_ := os.OpenFile(*filePath, os.O_RDWR |os.O_CREATE |os.O_APPEND,0644)
	f.Write([]byte(logStr))
	defer f.Close()
	fmt.Println("done...")

}