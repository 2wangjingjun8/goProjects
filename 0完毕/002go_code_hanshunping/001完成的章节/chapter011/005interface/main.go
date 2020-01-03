package main

import (
	"fmt"
)

type Ainterface interface{
	say()
}

type interger int
func (i interger) say()  {
	fmt.Println(i)
}


type stu struct{

}
func (i stu) say()  {
	fmt.Println("stu say()")
}

func main(){
	var i interger = 10
	var a Ainterface = i
	a.say()

	var stu stu
	var b Ainterface = stu
	b.say()
}