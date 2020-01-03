package main

import "fmt"

func main() {
	var ch chan int
	if ch == nil {
		fmt.Println("ch不能使用，需要进行make才能使用")
		ch = make(chan int)
		fmt.Printf("ch %T ch = %v", ch, ch)
	}
}
