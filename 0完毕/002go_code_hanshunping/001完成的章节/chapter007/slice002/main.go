package main

import (
	"fmt"
)

func test(slice []int)  {
	slice[1] = 20
}

func main()  {
	// var arr [5]int = [...]int{1,2,3,4,5}
	// slice1 := arr[:]
	// slice2 := slice1
	// slice2[1] = 10
	// fmt.Println("arr = ",arr)
	// fmt.Println("slice1 = ",slice1)
	// fmt.Println("slice2 = ",slice2)

	var slice = []int{1,2,3,4}
	fmt.Println("slice = ",slice)
	test(slice)
	fmt.Println("slice = ",slice)
}