package main

import (
	"fmt"
)
type Usb interface{
	start()
	stop()
}

type Phone struct{

}
func (p Phone) start(){
	fmt.Println("手机正在开始工作。。。")
}
func (p Phone) stop(){
	fmt.Println("手机已经结束工作。。。")
}

type Camera struct{

}
func (c Camera) start(){
	fmt.Println("相机正在开始工作。。。")
}
func (c Camera) stop(){
	fmt.Println("相机已经结束工作。。。")
}

type Computer struct{

}
func (computer Computer) working(usb Usb){
	usb.start()
	usb.stop()
}

func main()  {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	computer.working(phone)
	computer.working(camera)
	// fmt.Println()
}