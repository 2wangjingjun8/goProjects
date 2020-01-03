package main

import(
	"fmt"
)

type Student struct{
	name string
	gender string
	age int
	id int
	score float64
}

type monster struct{
	name string
	age int
}

func (stu *Student) say() string{
	res := fmt.Sprintf("学生\t姓名：%v,性别：%v,年龄：%v,学号：%v,分数：%v",stu.name,stu.gender,
	stu.age,stu.id,stu.score)
	return  res
}



func main()  {
	// var stu Student
	// stu.name= "小王"
	// stu.gender = "男"
	// stu.age = 20
	// stu.id = 2014096023
	// stu.score = 569.0
	// res := stu.say()
	// fmt.Println(res)

	var monster1 = monster{"小白",200}
	monster2 := monster{
		name:"小青",
		age:400,
	}


	var monster3 = &monster{"小虎",500}
	monster4 := &monster{
		name:"小狐",
		age:460,
	}
	fmt.Println(monster1,monster2,*monster3,*monster4)
}