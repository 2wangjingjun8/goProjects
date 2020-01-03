package main

import "fmt"

func main() {
	ch1 := make(chan string)   // 可读可写
	ch2 := make(chan<- string) // 只写
	ch3 := make(<-chan string) // 只读

	go write(ch1)
	go write(ch2)
	go read(ch3)
	read(ch1)
	// fmt.Println("ch1 = ", <-ch1)

}
func write(ch chan<- string) {
	ch <- "hello,I'm xiaoxiao"
}
func read(ch <-chan string) {
	data := <-ch
	fmt.Println("data = ", data)
}
