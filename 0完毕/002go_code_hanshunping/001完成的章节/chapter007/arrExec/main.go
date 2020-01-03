package main

import (
	"fmt"
)

func main()  {
	
	// arr := [26]byte{}
	var arr [26]byte
	arr[0] = 'A'
	for index := 0; index < len(arr); index++ {
		if index > 0 {
			arr[index] = arr[index-1] + 1
		}
		fmt.Printf("%c ",arr[index])
	}
	fmt.Println("")

	var arr2 = [...]int{10,-2,6,5,90,100,5}
	maxValue := arr2[0]
	maxindex := 0
	for i,v :=  range arr2{
		if maxValue < v {
			maxValue = v
			maxindex = i
		}
	}
	fmt.Printf("maxValue = %v,  maxindex = %v\n\n",maxValue, maxindex)

	var arr3 = [...]int{10,-2,6,5,90,100,5}
	sum := 0
	for _,v :=  range arr3{
		sum += v
	}
	fmt.Printf("sum = %v,  avg = %v\n",sum, float64(sum)/float64(len(arr3)))


}