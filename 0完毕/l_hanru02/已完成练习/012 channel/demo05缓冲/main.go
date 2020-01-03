package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := make(chan string, 4)
	go sendData(ch)
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("读取完毕...")
			break
		}
		fmt.Printf("ch: %v %v\n", len(ch), cap(ch))
		fmt.Printf("读取的数据是%v %v\n", v, ok)
	}

}

func sendData(ch chan string) {
	for i := 0; i <= 10; i++ {
		ch <- "数据" + strconv.Itoa(i)
		fmt.Printf("\t写入第%v个数据\n", i)
	}
	close(ch)

}
