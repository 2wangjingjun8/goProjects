package main

import (
	"fmt"
	"time"
)

func main() {
	// now := time.Now()
	// fmt.Println("now = ", now)
	// timer := time.NewTimer(3 * time.Second)
	// fmt.Println("<-timer.C = ", <-timer.C)
	execTime2()
}

func execTime() {
	fmt.Println("now = ", time.Now())
	timer := time.NewTimer(5 * time.Second)
	go func() {
		fmt.Println(<-timer.C)
		fmt.Println("结束。。。")
	}()
	time.Sleep(3 * time.Second)
	timer.Stop()
	fmt.Println("程序结束。。。")
}

func execTime2() {
	fmt.Println("now = ", time.Now())
	c := time.After(1 * time.Second)
	go func() {
		fmt.Println(<-c)
		fmt.Println("结束。。。")
	}()
	time.Sleep(3 * time.Second)
	fmt.Println("程序结束。。。")
}
