package main
import (
	"fmt"
)

type Circle struct {
	radius float64 
}

func (a Circle) area() float64 {
	return 3.14 * a.radius * a.radius
}

func main()  {
	var p Circle
	p.radius = 20
	res := p.area()

	fmt.Println("res = ",res)
}