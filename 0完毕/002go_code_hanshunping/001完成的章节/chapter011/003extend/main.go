package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age int
	Score float64
}
func (p *Student) ShowInfo(){
	fmt.Printf("名字：%v,年龄：%v,成绩：%v\n",p.Name,p.Age,p.Score)
}
func (p *Student) SetScore(score float64){
	p.Score = score
}

type Pupil struct {
	Student
}
func (p *Pupil) test(){
	fmt.Println("小学生正在考试。。。")
}


type Graduate struct {
	Student
}
func (p *Graduate) test(){
	fmt.Println("大学生正在考试。。。")
}

type Goods struct{
	Name string
	Price float64
}
type Brand struct{
	Name string
	Address string
}

type Tv struct{
	*Goods
	*Brand
}


func main()  {
	var tv = Tv{&Goods{"gd01",20.00},&Brand{"天猫","阿里巴巴"}}
	fmt.Println(*tv.Goods)
	// p := &Pupil{}
	// p.Name = "xiaoxiao"
	// p.Student.Age = 15
	// p.Student.SetScore(70.0)
	// p.Student.ShowInfo()
	// p.test()
	
	// g := &Graduate{}
	// g.Name = "xiaoxiao"
	// g.Student.Age = 20
	// g.Student.SetScore(90.0)
	// g.Student.ShowInfo()
	// g.test()
	
}