package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 获取当前时间
	t1 := time.Now()
	fmt.Printf("t1 %T = %v\n", t1, t1)

	// 获取指定时间
	t2 := time.Date(2018, 12, 20, 12, 05, 06, 0, time.Local)
	fmt.Printf("t2 = %v\n", t2)

	// time->string
	s1 := t1.Format("2006年1月2日 15:04:05\n")
	fmt.Printf("s1 = %v", s1)

	// string->time
	s2 := "2009年5月6日"
	t2, err := time.Parse("2006年1月2日", s2)
	if err != nil {
		fmt.Println("err = ", err)
	}
	fmt.Println(t2)

	y, m, d := t1.Date()
	fmt.Printf("y = %v\n", y)
	fmt.Printf("m = %v\n", m)
	fmt.Printf("d = %v\n", d)

	// 时间戳
	tNow := time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC)
	timestamp1 := tNow.Unix()
	fmt.Printf("timestamp1 = %v\n", timestamp1)
	timestamp2 := t1.Unix()
	fmt.Printf("timestamp2 = %v\n", timestamp2)

	timestamp3 := tNow.UnixNano()
	fmt.Printf("timestamp3 = %v\n", timestamp3)
	timestamp4 := t1.UnixNano()
	fmt.Printf("timestamp4 = %v\n", timestamp4)

	fmt.Printf("t1 = %v\n", t1)
	t5 := t1.Add(time.Second)
	fmt.Printf("t5 = %v\n", t5)
	t6 := t1.AddDate(1, 0, 0)
	fmt.Printf("t6 = %v\n", t6)

	// 随机数
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10) + 1
	fmt.Printf("n = %v\n", n)
	fmt.Println("睡眠1秒")
	time.Sleep(1 * time.Second)
	fmt.Println("睡醒了")
}
