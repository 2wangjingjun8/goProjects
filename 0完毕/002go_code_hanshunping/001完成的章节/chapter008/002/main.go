package main

import (
	"fmt"
)

func erfenFind(arr *[10]int,startIndex int,endIndex int,value int)  {
	if startIndex > endIndex {
		fmt.Println("找不到了")
	}
	middle := (startIndex + endIndex) / 2
	if (*arr)[middle] > value {
		erfenFind(arr,startIndex,middle - 1,value)
	}else if (*arr)[middle] < value {
		erfenFind(arr,middle + 1,endIndex,value)
	}else if (*arr)[middle] == value{
		fmt.Println("找到了，下标为", middle)
	}
	
}

func main()  {
	var arr = [10]int{1,5,10,45,90,110,452,550,1000,1200}
	erfenFind(&arr,0,len(arr),10)
	fmt.Println(arr)
}