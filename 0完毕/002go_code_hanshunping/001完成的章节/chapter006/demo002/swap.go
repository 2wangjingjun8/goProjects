package main

import (
	"fmt"
)
func swap(n1 *int, n2 *int) {
	temp := *n1
	*n1 = *n2
	*n2 = temp
}

func sumAll(args... int) int {
	sum := 0;
	for i := 0; i < len(args); i++{
		sum += args[i]
	}
	return sum
}

func main() {
	n1 :=  10
	n2 := 20
	swap(&n1,&n2)
	fmt.Println("n1 = ",n1,"n2 = ",n2)
	sum := sumAll(20,10,40,80,10)
	fmt.Println("sum = ",sum)
}