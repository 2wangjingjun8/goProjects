package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var num1 int = 99
	var num2 float64 = 10.235
	var b bool = true
	var char byte = 'h'
	var str string
	str = fmt.Sprintf("%d",num1)
	fmt.Printf("str type %T, str =  %q\n",str,str)

	str = fmt.Sprintf("%f",num2)
	fmt.Printf("str type %T, str =  %q\n",str,str)

	str = fmt.Sprintf("%t",b)
	fmt.Printf("str type %T, str =  %q\n",str,str)

	str = fmt.Sprintf("%c",char)
	fmt.Printf("str type %T, str =  %q\n",str,str)

	str = strconv.FormatInt(int64(num1),10)
	fmt.Printf("str type %T, str =  %q\n",str,str)
	str = strconv.FormatFloat(num2,'f',10,64)
	fmt.Printf("str type %T, str =  %q\n",str,str)
	str = strconv.FormatBool(b)
	fmt.Printf("str type %T, str =  %q\n",str,str)

}