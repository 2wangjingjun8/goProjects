package main

import (
	"fmt"
)

func main() {
	var slice []int = make([]int,4,8)
	slice[1] = 10
	slice[3] = 20
	fmt.Println(slice)
	fmt.Printf("slice的容量：%v",cap(slice))

	var slicestr []string = []string{"maomao","xiaoxiao","fada","fada"}
	fmt.Println(slicestr)
	fmt.Println("slicestr size = ",len(slicestr))
	fmt.Println("slicestr cap = ",cap(slicestr))


}