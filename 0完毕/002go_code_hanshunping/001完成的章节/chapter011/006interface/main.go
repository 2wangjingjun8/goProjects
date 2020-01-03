package main

import (
	"fmt"
)

type Usb interface{
	say()
}

type Stu struct {

}
func (s *Stu) say(){
	fmt.Println("say()")

}

func main(){
	var stu Stu
	var a Usb = &stu
	a.say()

	
}