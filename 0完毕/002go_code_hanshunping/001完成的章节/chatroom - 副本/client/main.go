package main

import(
	"fmt"
)
var (
	userId int
	pwd string
)

func main()  {
	var loop = true
	var key int
	for loop {
		fmt.Println("----------欢迎登陆多人聊天系统----------")
		fmt.Println("----------1 登陆")
		fmt.Println("----------2 注册")
		fmt.Println("----------3 退出")
		fmt.Println("----------请选择（1-3）")
		fmt.Scanf("%d\n",&key)
		switch key {
			case 1:
				fmt.Println("----------登陆----------")
				loop = false
			case 2:
				fmt.Println("----------注册----------")
				loop = false
			case 3:
				fmt.Println("----------退出----------")
				loop = false
			default:
				fmt.Println("输入有误，请重新输入。。。")
		}
	}

	if key == 1 {
		// ----------登陆----------
		fmt.Println("请输入id")
		fmt.Scanf("%d\n",&userId)
		fmt.Println("请输入密码")
		fmt.Scanf("%s\n",&pwd)
		login(userId, pwd)
		
	}else if key == 2 {
		fmt.Println("----------注册----------")
	}

}