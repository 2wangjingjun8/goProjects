package main

import (
	"fmt"
	"go_code/chapter011/002student/student"
)


func main()  {
	stu := student.NewStudent("xiaoxaio",20,200.0)
	var age = stu.GetAge()
	fmt.Println(age)
	stu.SetAge(18)
	age = stu.GetAge()
	fmt.Println(age)

}