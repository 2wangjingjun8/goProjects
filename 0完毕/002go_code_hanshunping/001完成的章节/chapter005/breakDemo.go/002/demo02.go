package main

import (
	"fmt"
)
func main() {
	num := 3
	name := ""
	password := ""

	for i:=0;i<3;i++ {
		fmt.Println("请输入用户名：")
		fmt.Scanln(&name)
		fmt.Println("请输入密码：")
		fmt.Scanln(&password)
		if name == "小小" && password == "123456" {
			fmt.Println("登录成功")
			break
		}else{
			num--
			fmt.Printf("亲，账号或密码错误，还剩下%d次机会\n",num)
			
		}
	}
	if num == 0 {
		fmt.Printf("机会用完")
	}
}