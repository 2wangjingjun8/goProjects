package main

import (
	"fmt"
)
func addAndSub(a int, b int) (sum int, sub int)  {
	sum = a + b
	sub = a - b
	return 
}

func main()  {
	sum,sub := addAndSub(10,20)
	fmt.Println("sum = ", sum, "sub = ", sub)
}