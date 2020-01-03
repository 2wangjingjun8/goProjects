package utils

import(
	"fmt"
)

type FamilyAccount struct{
	details string
	balance float64
	money float64
	note string
	flag bool
	key string
	loop bool
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		details:"收支\t\t账户余额\t收支金额\t收支说明\n",
		balance:10000.00,
		money:0.00,
		note:"",
		flag:false,
		key:"",
		loop:true,
	}
}

func (this *FamilyAccount) detail()  {
	fmt.Println("------------收支明细------------")
	if this.flag {
		fmt.Println(this.details)
	}else{
		fmt.Println("当前没有收支记录，来一笔收支吧...")
	}
}
func (this *FamilyAccount) income()  {
	// fmt.Println("------------登记收入------------")
	fmt.Printf("请输入收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Printf("请输入收入说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("收入\t\t%v\t\t%v\t\t%v\n",this.balance,this.money,this.note)
	this.flag = true
}
func (this *FamilyAccount) pay()  {
	// fmt.Println("------------登记支出------------")
	fmt.Printf("请输入支出金额：")
	fmt.Scanln(&this.money)
	fmt.Printf("请输入支出说明：")
	fmt.Scanln(&this.note)
	if this.balance > this.money {
		this.balance -= this.money
	}else{
		fmt.Println("余额不足")
	}
	this.details += fmt.Sprintf("支出\t\t%v\t\t%v\t\t%v\n",this.balance,this.money,this.note)
	this.flag = true
}
func (this *FamilyAccount) exit()  {
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
		this.loop = false
	}
}

func (this *FamilyAccount) ShowMenuMain()  {
	for{
		fmt.Println("------------家庭收支记账软件------------")
		fmt.Println("            1 收支明细")
		fmt.Println("            2 登记收入")
		fmt.Println("            3 登记支出")
		fmt.Println("            4 退出软件")
		fmt.Printf("请选择（1-4）：")
		fmt.Scanln(&this.key)

		switch this.key {
			case "1":
				this.detail()
			case "2":
				this.income()
			case "3":
				this.pay()
			case "4":
				this.exit()
			default:
				fmt.Println("您输入的选择有误...")
		}
		if !this.loop {
			break
		}
	}
}