package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 100
	}()
	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "ahha"
	}()
	select {
	case v, ok := <-ch1:
		if ok {
			fmt.Println("ch1执行...v= ", v)
		} else {
			fmt.Println("ch1通道结束...")
		}
	case <-ch2:
		fmt.Println("ch2执行...")
	case <-time.After(3 * time.Second):
		fmt.Println("time.After执行。。。")
		// default:
		// 	fmt.Println("default执行...")
	}
}
