package main

import "fmt"

func febonacci() func() int {
	var a, b = 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
func main() {
	f := febonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
