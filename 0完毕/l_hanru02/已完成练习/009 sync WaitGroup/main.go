package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func1()
	go func2()
	fmt.Println("主进程进入阻塞状态...")
	wg.Wait()
	fmt.Println("主进程阻塞状态解除...")
	fmt.Println("程序结束！！！")
}

// 打印字母
func func1() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("打印第%v个字母\n", i)
	}
	wg.Done()

}

// 打印数字
func func2() {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		fmt.Printf("\t打印第%v个数字\n", i)
		// runtime.Gosched() // 让出时间片给func1先执行
	}

}
