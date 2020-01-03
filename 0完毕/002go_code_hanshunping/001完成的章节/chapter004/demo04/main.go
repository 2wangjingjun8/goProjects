package main

import (
	"fmt"
)

func main()  {
	var a = 20
	var b = 30
	fmt.Printf("改变之前: a = %d, b = %d\n",a,b)
	a,b = b,a
	fmt.Printf("改变之后: a = %d, b = %d",a,b)
}