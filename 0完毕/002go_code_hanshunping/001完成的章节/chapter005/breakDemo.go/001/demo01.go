package main

import (
	"fmt"
)
func main()  {
	var sum = 0
	for i:=1;i<100;i++ {
		sum += i
		fmt.Println(i)
		if sum > 20{
			fmt.Println(sum)
			break
		}
	}
}