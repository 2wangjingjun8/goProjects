package main

import "fmt"

func insertSort(arr []int) {
	fmt.Println(arr)
	for i := range arr {
		if i == 0 {
			continue
		}
		fmt.Println(i)
		preIndex := i - 1
		current := arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex+1] = current
		fmt.Println(arr)
	}
}
func main() {
	var arr = []int{10, 6, 8, 2, 20, -1, -20}
	insertSort(arr)
}
