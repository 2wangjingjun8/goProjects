package main

import "fmt"

func main()  {
	var i = 0
	for {
		fmt.Println("hello world")
		if i>10 {
			break
		}
		i++
	}
	
}