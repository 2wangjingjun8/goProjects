package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go sendData(ch)

	// for {
	// 	if data, ok := <-ch; ok {
	// 		fmt.Printf("取出data = %v\n", data)
	// 	} else {
	// 		fmt.Println("通道结束...")
	// 		break
	// 	}
	// }
	for val := range ch {
		fmt.Printf("取出data = %v\n", val)

	}

	fmt.Println("程序运行结束...")
}

func sendData(ch chan int) {
	for i := 0; i <= 10; i++ {
		// time.Sleep(500 * time.Millisecond)
		fmt.Printf("%v被写入通道\n", i)
		ch <- i
	}
	close(ch)
}
