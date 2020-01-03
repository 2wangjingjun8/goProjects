package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		// fmt.Println("子routine开始执行...")
		fmt.Println("data = ", <-ch)
	}()
	ch <- 10
	fmt.Println("程序执行完毕")
}
