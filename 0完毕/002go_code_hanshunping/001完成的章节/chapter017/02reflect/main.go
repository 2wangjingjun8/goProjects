package main

import(
	"fmt"
	"reflect"
)

type Student struct{
	Name string
	Age int
}

func reflect01(b interface{})  {
	rVal := reflect.ValueOf(b)
	rVal.Elem().SetInt(50)
}

func main()  {
	stu := Student{"haha",20}
	var n int = 20
	reflect01(&n)
	fmt.Println(stu)
	fmt.Println("n = ",n)
	
}