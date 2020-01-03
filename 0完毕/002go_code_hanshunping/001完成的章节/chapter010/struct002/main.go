package main
import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func main()  {
	// 第一种声明
	// var person Person

	// 第二种声明
	var person = Person{}
	person.Name = "maomao"
	person.Age = 156
	fmt.Println(person)

	// 第三种方式
	p3 := new(Person)
	(*p3).Name = "haha"
	fmt.Println(*p3)

	// 第四种方式
	p4 := &Person{}
	p4.Name = "maka"
	p4.Age = 230
	fmt.Println(*p4)

}