package main

import(
	"fmt"
)

func main()  {
	var aa chan int
	aa = make(chan int, 3)
	aa<-99
	num := 10
	aa <- num
	fmt.Printf("aa = %v\n",aa)
	b1 := <-aa
	b2 := <-aa
	b3 := <-aa
	fmt.Printf("b1 = %v\n",b1)
	fmt.Printf("b2 = %v\n",b2)
	fmt.Printf("b3 = %v\n",b3)
}