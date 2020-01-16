package main

import (
	"fmt"
	_ "test/utils"
)

func showNum(n1 int) {
	fmt.Println(n1)
}

func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	return sum
}

func main() {
	// fmt.Println(u.Name)
	// fmt.Println(u.GetName())
	n1 := 20
	showNum(n1)
	sum := getSum(20,10)
	fmt.Println(sum)

}