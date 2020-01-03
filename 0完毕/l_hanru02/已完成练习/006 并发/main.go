package main

import (
	"fmt"
	"time"
)

func main() {
	go printNum()
	for i := 0; i <= 1000; i++ {
		fmt.Printf("打印第%d个字母\n", i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Exec over...")
}

func printNum() {
	for i := 0; i <= 1000; i++ {
		fmt.Printf("打印第%d个数字\n", i)
	}
}
