package main

import (
	"fmt"
)

type Account struct{
	AccountNo string
	Pwd string
	Balance float64
}

func (a *Account) SaveMoney(num float64,pwd string){
	if a.Pwd != pwd {
		fmt.Println("密码错误，请重新输入")
		return
	}
	a.Balance += num
	fmt.Printf("存入%v元，您的账户余额为：%v\n",num,a.Balance)

}

func (a *Account) GetMoney(num float64,pwd string){
	if a.Pwd != pwd {
		fmt.Println("密码错误，请重新输入")
		return
	}
	a.Balance -= num
	fmt.Printf("取走%v元，您的账户余额为：%v\n",num,a.Balance)

}

func (a *Account) QueryMoney(pwd string){
	if a.Pwd != pwd {
		fmt.Println("密码错误，请重新输入")
		return
	}
	fmt.Printf("您的账户余额为：%v\n",a.Balance)

}

func main() {
	p1 := Account{"gs00110","123456",100.00}
	p1.QueryMoney("123456")	
	p1.SaveMoney(200.0,"123456")	
	p1.GetMoney(150.0,"123456")	
}