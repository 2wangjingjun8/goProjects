package main

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	// 设置go程序执行的最大cpu数量【1，256】
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("n=", n)
}

func main() {
	// fmt.Println("goroot->", runtime.GOROOT())
	// fmt.Println("goos->", runtime.GOOS)
	// fmt.Println("NumCPU->", runtime.NumCPU())

	//runtime.Gosched 让出时间片，先让其他goroutine执行
	// go func() {
	// 	for i := 0; i < 5; i++ {
	// 		fmt.Println("goroutine i = ", i)
	// 	}
	// }()

	// for i := 0; i < 5; i++ {
	// 	// 让出时间片，先让其他goroutine执行
	// 	runtime.Gosched()
	// 	fmt.Println("main i = ", i)

	// }
	go func() {
		fmt.Println("goroutine 开始")
		fun()
		fmt.Println("goroutine 结束")
	}()
	time.Sleep(2 * time.Second)
}
func fun() {
	defer fmt.Println("defer...")
	runtime.Goexit()
	fmt.Println("fun...")
}
