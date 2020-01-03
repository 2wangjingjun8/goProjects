package main

import(
	"fmt"
)

func insertSort(arr *[5]int)  {
	// 第一个不动，从第二个开始
	for i := 1; i < len(arr); i++ {
		curVal := arr[i] 
		insertIndex := i-1

		for insertIndex >= 0 && arr[insertIndex] < insertIndex {
			arr[insertIndex] = arr[insertIndex+1]
			insertIndex--
		}

		if insertIndex +1 != i {
			arr[insertIndex+1] = curVal
		}
	}
}

func main()  {
	arr := [5]int{10,30,5,40,15}
	fmt.Println("排序前：arr = ",arr)
	insertSort(&arr)
	fmt.Println("排序后：arr = ",arr)


}