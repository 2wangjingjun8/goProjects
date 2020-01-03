package main
import(
	"fmt"
)

func selectSort(arr *[6]int)  {
	for j := 0; j < len(arr)-1; j++ {
		maxIndex := j
		max := arr[j]
		for i := j; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		// 交换
		arr[j],arr[maxIndex] = arr[maxIndex],arr[j]
		fmt.Println(arr)
	}
}

func main()  {
	arr := [6]int{4,2,3,5,1,6}
	fmt.Println(arr)
	selectSort(&arr)
}