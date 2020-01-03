package main

import (
	"fmt"
)
type stu struct {
	name string
	age int
	address string
}

func main()  {
	var map1 = make(map[string]stu, 0)
	stu1 := stu{"xiaoxiao",18,"广州"}
	stu2 := stu{"maoamo",19,"深圳"}
	map1["stu1"] = stu1
	map1["stu2"] = stu2
	fmt.Println(map1)

	for i,v := range map1{
		fmt.Printf("学号为%v的学生，姓名为%v\n",i,v.name)
		fmt.Printf("学号为%v的学生，年龄为%v\n",i,v.age)
		fmt.Printf("学号为%v的学生，地址为%v\n\n",i,v.address)
	}
}