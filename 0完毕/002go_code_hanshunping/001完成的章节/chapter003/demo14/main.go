package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var str string = "true"
	var num int64
	var b bool
	var str1 string= "123.201"
	var f float64
	num,_ = strconv.ParseInt(str,10,64)
	fmt.Printf("num type %T, num =  %v\n",num,num)
	b,_ = strconv.ParseBool(str)
	fmt.Printf("b type %T, b =  %v\n",b,b)
	f,_ = strconv.ParseFloat(str1,64)
	fmt.Printf("f type %T, f =  %v\n",f,f)

}