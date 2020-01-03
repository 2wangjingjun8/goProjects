package main

import(
	"fmt"
	_"time"
)



func main()  {
	// 只读
	// var chan1 <-chan int
	// num1 := <-chan1
	// fmt.Println(num1)

	// 只写
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2<-20
	// num2 := <-chan2
	fmt.Println(chan2)
}