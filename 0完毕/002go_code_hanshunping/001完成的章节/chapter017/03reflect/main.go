package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var str string = "hello world"
	rVal := reflect.ValueOf(&str)
	rVal.Elem().SetString("Jack, hello world")
	fmt.Printf("rVal = %v,type = %T\n",rVal,rVal)
	fmt.Printf("str = %v,type = %T\n",str,str)
}