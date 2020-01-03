package main

import "fmt"

func sum(args ...int) {
	sum := 0
	for _, v := range args {
		sum += v
	}
	fmt.Println("sum = ", sum)
}

func main() {
	sum(1, 2, 3, 4, 5, 6)
}
