package main

import(
	"fmt"
)

func main()  {

	var details = "收支\t\t账户余额\t收支金额\t收支说明\n"
	//总金额
	balance := 10000.00
	//收支金额
	money := 0.00
	//收支说明
	note := ""
	//收支有没有记录的标志
	flag := false
	
	var key = ""
	var loop = true


	for{
		fmt.Println("------------家庭收支记账软件------------")
		fmt.Println("            1 收支明细")
		fmt.Println("            2 登记收入")
		fmt.Println("            3 登记支出")
		fmt.Println("            4 退出软件")
		fmt.Printf("请选择（1-4）：")
		fmt.Scanln(&key)

		switch key {
			case "1":
				fmt.Println("------------收支明细------------")
				if flag {
					fmt.Println(details)
				}else{
					fmt.Println("当前没有收支记录，来一笔收支吧...")
				}
				
			case "2":
				// fmt.Println("------------登记收入------------")
				fmt.Printf("请输入收入金额：")
				fmt.Scanln(&money)
				balance += money
				fmt.Printf("请输入收入说明：")
				fmt.Scanln(&note)
				details += fmt.Sprintf("收入\t\t%v\t\t%v\t\t%v\n",balance,money,note)
				flag = true

			case "3":
				// fmt.Println("------------登记支出------------")
				fmt.Printf("请输入支出金额：")
				fmt.Scanln(&money)
				fmt.Printf("请输入支出说明：")
				fmt.Scanln(&note)
				if balance > money {
					balance -= money
				}else{
					fmt.Println("余额不足")
				}
				details += fmt.Sprintf("支出\t\t%v\t\t%v\t\t%v\n",balance,money,note)
				flag = true
			case "4":
				// fmt.Println("------------退出软件------------")
				var choice = ""
				for{
					fmt.Printf("确定要退出软件吗？{y/n}")
					fmt.Scanln(&choice)
					if choice == "y" || choice == "n" {
						break
					}
					fmt.Printf("请输入正确的命令\n")
				}
				if choice == "y" {
					loop = false
				}
			default:
				fmt.Println("您输入的选择有误...")
		}
		if !loop {
			break
		}


	}
	fmt.Println("您已退出软件...")

}