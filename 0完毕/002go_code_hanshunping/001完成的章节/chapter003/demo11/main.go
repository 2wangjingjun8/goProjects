package main

import (
	"fmt"
)

func main()  {
	var a1 int
	var a2 float32
	var a3 float64
	var a4 bool
	var a5 string

	fmt.Printf("a1 = %d,a2 = %v,a3 = %v,a4 = %v,a5 = %v",a1,a2,a3,a4,a5)

	var i int32 = 100

	var n1 float64 = float64(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)
	fmt.Printf("n1 = %v,n2 = %v,n3 = %v",n1,n2,n3)






}