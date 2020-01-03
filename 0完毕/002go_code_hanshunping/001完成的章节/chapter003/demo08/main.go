package main

import (
	"fmt"
)

func main()  {
	var a float32 = 0.00096541124512
	var b float64 = 0.00096541124512
	fmt.Println("a = ",a,"b = ",b)

	var c = 0.0278
	fmt.Printf("c的类型是%T\n",c)

	var d = 0.2345
	var e = .2345
	fmt.Println("d = ",d,"e = ",e)

	var f = 5.2073e2
	fmt.Println("f = ",f)
	var g = 5.2073E2
	fmt.Println("g = ",g)
	
	var h = 5.2073e-2
	fmt.Println("h = ",h)


}