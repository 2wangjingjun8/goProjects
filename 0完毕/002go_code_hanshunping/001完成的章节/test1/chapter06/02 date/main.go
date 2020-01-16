package main

import(
	"fmt"
	"time"
	"strconv"
)

func test() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello"+strconv.Itoa(i)
	}
}

func main() {
	now := time.Now()
	fmt.Printf("now:%v\n类型是%T\n",now,now)

	// // fmt.Printf("当前时间为：%d-%d-%d %d:%d:%d",
	// 	now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	// a := fmt.Sprintf("当前时间为：%d-%d-%d %d:%d:%d",
	// 	now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	// fmt.Println(a)

	fmt.Printf("时间：%v",now.Format("2006-01-02 15:04:05"))
	// fmt.Printf("时间年：%v",now.Format("15"))

	// i := 0
	// for{
	// 	i++
	// 	fmt.Println(i)
	// 	time.Sleep(time.Millisecond*100)
	// 	if i>100 {
	// 		break
	// 	}
	// }
	fmt.Printf("unix:%v,unixnano:%v",now.Unix(),now.UnixNano())
	start := time.Now().Unix()
	// test()
	end := time.Now().Unix()
	fmt.Printf("执行的时间是%v秒",end-start)



}