package main

import (
	"fmt"
)

type Monkey struct{
	Name string
}
func (this Monkey) Climbing(){
	fmt.Printf("%v生来就能爬树\n",this.Name)
}

type LittleMonkey struct{
	Monkey
}
func (this *LittleMonkey) flying(){
	fmt.Printf("%v学会了飞翔\n",this.Name)
}
func (this *LittleMonkey) Swimmng(){
	fmt.Printf("%v学会了游泳\n",this.Name)
}


type LLMonkey struct{
	LittleMonkey
}

type Bird interface{
	flying()
}
type Fish interface{
	Swimmng()
}

func main()  {
	monkey := LittleMonkey{
		Monkey{
			Name:"悟空",
		},
	}
	monkey.Climbing()
	monkey.flying()
	monkey.Swimmng()
	fmt.Println(0)
	LLMonkey := LLMonkey{}
	LLMonkey.Name = "孙者行"
	LLMonkey.flying()
}