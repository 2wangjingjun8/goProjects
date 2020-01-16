package main

import "fmt"

func erfenFind(arr *[7]int,startIndex int ,endIndex int,value int) {
	fmt.Println(startIndex,endIndex)
	if startIndex >= endIndex {
		fmt.Println("找不到")
		return
	}
	middle := (startIndex + endIndex)/2
	fmt.Println(middle)
	if (*arr)[middle] > value {
		erfenFind(arr,startIndex,middle-1,value)
	}else if (*arr)[middle] < value {
		erfenFind(arr,middle+1,endIndex,value)
	}else{
		fmt.Printf("找到了下标为%v",middle)
	}
}

func main() {
	var arr = [7]int{1,5,10,54,70,50,90}
	erfenFind(&arr,0 ,len(arr),90)
}