package main
import (
	"fmt"
)

type integer int

func (i integer) print()  {
	fmt.Println("i = ",i)
}

func (i *integer) change() {
	(*i) += 1 
}

type Students struct{
	Name string
	Age int
}

func (stu *Students) String() string {
	// fmt.Println(123)
	str := fmt.Sprintf("name = %v age = %v",stu.Name,stu.Age)
	return str
}

func main()  {
	// var p integer
	// p = 10
	// p.print()
	// p.change()
	// fmt.Println(p)

	stu := Students{"dada",20}
	fmt.Println(&stu)

}