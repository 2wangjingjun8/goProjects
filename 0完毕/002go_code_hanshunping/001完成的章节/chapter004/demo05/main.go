package main

import (
	"fmt"
)

func main()  {
	var name string 
	var age int
	var salary float32
	var isPass bool
	// fmt.Println("请输入名字：")
	// fmt.Scanln(&name)
	// fmt.Println("请输入年龄：")
	// fmt.Scanln(&age)
	// fmt.Println("请输入薪水：")
	// fmt.Scanln(&salary)
	// fmt.Println("请输入是否通过考试：")
	// fmt.Scanln(&isPass)
	fmt.Println("请输入名字 年龄 薪水 是否通过考试：（空格隔开）")
	fmt.Scanf("%v %d %f %t",&name,&age,&salary,&isPass)

	fmt.Printf("姓名：%v 年龄：%v 薪水：%v 通过与否：%v",name,age,salary,isPass)
}