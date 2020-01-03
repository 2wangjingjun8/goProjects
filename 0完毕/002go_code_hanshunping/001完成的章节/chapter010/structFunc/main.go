package main
import (
	"fmt"
)

type Person struct {
	Name string 
	Age int 
}

func (a Person) test(){
	fmt.Println("name = ",a.Name)
}

func main()  {
	var p = Person{}
	p.Name = "haha"
	p.Age = 18
	p.test()

	// fmt.Println()
}