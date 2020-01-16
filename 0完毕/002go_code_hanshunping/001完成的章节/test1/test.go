package main
import(
	"fmt"
)
func bubbleSort(arr []int){
	fmt.Println(arr)
	length := len(arr)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
	fmt.Println(arr)
}

func main()  {
	arr := []int{10,5,9,6,18,-5,0,4}
	bubbleSort(arr)
}