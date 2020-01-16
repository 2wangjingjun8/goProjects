package main

import "fmt"

type Point struct{
	x int
	y int
}
func main() {
	// var q interface{}
	// var f float32 = 1.1
	// q = f
	// bb := q
	// fmt.Println(q)
	// fmt.Println(bb)
	aa := Point{2,4}
	var in interface{}
	in = aa
	var y Point
	y = in.(Point)
	fmt.Println(in)
	fmt.Println(y)
	if in == y {
		fmt.Println("==")
	}else{
		fmt.Println("!=")
	}

}