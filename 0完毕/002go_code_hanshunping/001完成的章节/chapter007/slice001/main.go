package main

import (
	"fmt"
)

func main()  {
	var arr [5]int = [...]int{20,3,5,44,68}
	slice := arr[1:3]
	slice[1] = 10
	fmt.Println("arr的值：",arr)
	fmt.Println("slice的值：",slice)
	fmt.Println("slice的元素个数：",len(slice))
	fmt.Println("slice的容量：",cap(slice))
}