package main

import (
	"fmt"
)

func main()  {
	var second int
	fmt.Println("请输入秒数：")
	fmt.Scanln(&second)
	if second < 8 {
		var genter string
		fmt.Println("请输入性别：")
		fmt.Scanln(&genter)
		if genter == "男"{
			fmt.Println("进入男子决赛组")
		}else{
			fmt.Println("进入女子决赛组")
		}
	}else{
		fmt.Println("Out...")
	}

}