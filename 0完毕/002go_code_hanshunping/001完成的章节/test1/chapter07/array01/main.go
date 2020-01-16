package main

import (
	"fmt"
)

func main() {
	var arr [6]float64
	arr[0] = 8.0
	arr[1] = 4.0
	arr[2] = 5.0
	arr[3] = 1.0
	arr[4] = 2.0
	arr[5] = 15.0
	allWeight := 0.0
	for i := 0; i < len(arr); i++ {
		allWeight += arr[i]
	}
	avgWeight := fmt.Sprintf("%.2f",allWeight/float64(len(arr)))
	fmt.Printf("allWeight = %v , avgWeight = %v",allWeight,avgWeight)
}