package main

import (
	"fmt"
	"reflect"
)

func reInteface(b interface{})  {
	rType := reflect.TypeOf(b)
	fmt.Printf("rType = %v,type = %T\n",rType, rType)

	rVal := reflect.ValueOf(b)

	num1 := 2
	num2 := 10 + num1
	fmt.Printf("rVal = %v,type = %T \n num2 = %v\n",rVal, rVal, num2)

	num3 := rVal.Interface()
	num4 :=  num3.(int)
	fmt.Printf("num4 = %v,type = %T\n",num4, num4)
}


func reInteface2(b interface{})  {
	rType := reflect.TypeOf(b)
	fmt.Printf("rType = %v,type = %T\n",rType, rType)

	rVal := reflect.ValueOf(b)
	iV := rVal.Interface()

	kind1 := rVal.Kind()
	kind2 := rType.Kind()
	fmt.Printf("kind1 = %v,kind2 = %v\n",kind1, kind2)

	fmt.Printf("iV = %v,type = %T\n",iV, iV)
	// 类型断言
	stu := iV.(Student)
	fmt.Printf("Name = %v,type = %T\n",stu.Name, stu)
}

type Student struct{
	Name string
	Age int
}

func main()  {
	// var num = 20
	// reInteface(num)

	stu := Student{"xiuaowag",20}
	reInteface2(stu)
}