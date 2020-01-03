package student

import (
	_ "fmt"
)

type student struct{
	Name string
	age int
	sal float64
}

func NewStudent(name string, age int,sal float64) *student {
	return &student{
		Name:name,
		age:age,
		sal:sal,
	}
}

func (stu *student) SetAge(age int) {
	stu.age = age
}
func (stu *student) GetAge() int {
	return stu.age
}

func (stu *student) SetSal(sal float64) {
	stu.sal = sal
}
func (stu *student) GetSal() float64 {
	return stu.sal
}


