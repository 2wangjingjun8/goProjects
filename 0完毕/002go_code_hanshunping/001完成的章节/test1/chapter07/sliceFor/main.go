package main

import (
	"fmt"
)

func main() {
	var arr [5]int = [5]int{10,20,30,40,50}
	slice := arr[1:5]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v] = %v ",i,slice[i])
	}
	fmt.Println("")
	for i,v := range slice{
		fmt.Printf("i = %v v = %v \n",i,v)
	}

}