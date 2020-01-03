package main

import (
	"fmt"
)

func main()  {
	var arr [4][6]int
	for i := 0; i < len(arr); i++ {
		// fmt.Println(arr[i])
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v ",arr[i][j])
			
		}
		fmt.Println("")
	}
	
}