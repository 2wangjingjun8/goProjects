package main

import "fmt"

func main() {
	r := Add(1, 2)
	fmt.Println(r)
}

// Add is a function
func Add(a, b int) int {
	return a + b
}
