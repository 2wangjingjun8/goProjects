package main

import (
	"fmt"
)
type Usb interface{
	start()
	call()
	stop()
}

type Phone struct{
	Name string
}
func (p Phone) start(){
	fmt.Println("手机正在开始工作。。。")
}
func (p Phone) call(){
	fmt.Println("手机正在打电话。。。")
}
func (p Phone) stop(){
	fmt.Println("手机已经结束工作。。。")
}

type Camera struct{
	Name string
}
func (c Camera) start(){
	fmt.Println("相机正在开始工作。。。")
}
func (c Camera) stop(){
	fmt.Println("相机已经结束工作。。。")
}
func (c Camera) call(){
	fmt.Println("")
}

type Computer struct{

}
func (computer Computer) working(usb Usb){
	usb.start()
	if phone,ok := usb.(Phone);ok {
		phone.call()
	}
	usb.stop()
}

func main()  {
	computer := Computer{}
	var arr [3]Usb
	arr[0] =  Phone{"荣耀"}
	arr[1] =  Phone{"小米"}
	arr[2] =  Camera{"尼康"}

	for _,v := range arr{
		computer.working(v)
	}

	// phone := Phone{}
	// camera := Camera{}
	// computer.working(phone)
	// computer.working(camera)
	// fmt.Println()
}