package main

import "fmt"

var(
	n1 = 10
	n2 = 20
	name = "xiaoxiao"
)
func main()  {
	// 默认值
	// var i int
	// fmt.Printf("i = %d\n",i)

	// 不说明类型
	// var j = 10.11
	// fmt.Println(j )

	// 省略var, 用:=替换
	// m := 20
	// fmt.Println(m)

	// 多变量声明
	// var n1,n2,n3 int
	// fmt.Println( n1,n2,n3)
	
	// 多变量声明，（类型推导式
	// n1,n2,name := 0.11,20,"xiaoxiao"
	fmt.Println( n1,n2,name)


}