package main

import(
	"fmt"
)

type MethodUtils struct{
	num1 float64
	num2 float64
}

func (stu MethodUtils) operate(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = stu.num1 + stu.num2
	case '-':
		res = stu.num1 - stu.num2
	case '*':
		res = stu.num1 * stu.num2
	case '/':
		res =  stu.num1 / stu.num2
	default:
		fmt.Println("操作符错误")
	}
	return res
}


func main()  {
	stu := MethodUtils{20,10}
	res1 := stu.operate('+')
	fmt.Println("res1 = ",res1)
	res2 := stu.operate('-')
	fmt.Println("res2 = ",res2)
	
}