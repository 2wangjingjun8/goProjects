package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	var arr [5]int
	rand.Seed(time.Now().UnixNano())
	for index := 0; index < len(arr); index++ {
		arr[index] = rand.Intn(100)
	}
	fmt.Println(arr)
	var temp = 0
	for index := 0; index < len(arr)/2; index++ {
		temp = arr[len(arr)-1-index]
		arr[len(arr)-1-index] = arr[index]
		arr[index] = temp
	}
	fmt.Println(arr)
}