package main

import (
	"fmt"
)

func main()  {
	var days = 97
	var weeks = days / 7
	var day = days % 7
	fmt.Printf("weeks = %d days = %d\n",weeks,day)

	var f = 134.2
	var t = 5.0 / 9 * (f - 100)
	fmt.Printf("f = %v t = %v",f,t )


}