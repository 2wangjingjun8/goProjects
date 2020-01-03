package main

import (
	"fmt"
)

func arrSort(arr *[5]int) [5]int {
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j],(*arr)[j+1] = (*arr)[j+1],(*arr)[j]
			}
		}
		
	}
	return *arr

}

func main()  {
	var arr = [5]int{10,90,50,60,16}
	aa := arrSort(&arr)
	fmt.Println(aa)
}