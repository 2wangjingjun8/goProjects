package main

import (
	"fmt"
)

func main() {
	var ch chan bool
	ch = make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("打印第%v个数字\n", i)
		}
		fmt.Println("打印结束")
		ch <- true
	}()
	fmt.Println("准备读取data")
	data := <-ch
	fmt.Println("data -->", data)
	fmt.Println("程序执行结束")
}
