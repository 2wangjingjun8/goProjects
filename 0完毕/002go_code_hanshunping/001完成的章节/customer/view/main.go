package main

import (
	"fmt"
	"go_code/customer/service"
	"go_code/customer/model"
)

type customerView struct{
	key string
	loop bool
	customerService *service.CustomerService

}

func (this *customerView) list() {
	fmt.Println("----------------客户列表----------------")
	customers := this.customerService.List()
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		// 获取一个人的字符串信息
		fmt.Println(customers[i].GetInfo())
		
	}
	fmt.Println("----------------客户列表完成----------------")
}

func (this *customerView) add(){
	fmt.Println("----------------添加客户----------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	genter := ""
	fmt.Scanln(&genter)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer2(name,genter,age,phone,email)
	this.customerService.Add(customer)
	fmt.Println("----------------添加客户完成----------------")
}

func (this *customerView) delete(){
	fmt.Println("----------------删除客户----------------")
	fmt.Printf("请输入要删除客户的编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1{
		return
	}
	choice := ""
	fmt.Printf("确定要删除吗？（y/n）：")
	for{
		fmt.Scanln(&choice)
		if choice =="y" || choice == "Y" {
			if this.customerService.Delete(id) {
				fmt.Println("----------------删除完成----------------")
				return
			}else{
				fmt.Println("----------------删除失败，输入的编号不存在----------------")
				return
			}
		}else if(choice =="n" || choice == "N"){
			return
		}else{
			fmt.Printf("输入错误，请重新输入（y/n）：")
		}
	}
}


func (this *customerView) update(){
	fmt.Println("----------------删除客户----------------")
	fmt.Printf("请输入要修改客户的编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1{
		return
	}
	var customer = this.customerService.FindOneByID(id)
	
	// if customer {
	// 	fmt.Println("找不到该编号的客户")
	// 	return
	// }
	fmt.Printf("姓名(%v)：",customer.Name)
	name := ""
	fmt.Scanln(&name)
	if name == ""{
		name = customer.Name
	}

	fmt.Printf("性别(%v)：",customer.Genter)
	genter := ""
	fmt.Scanln(&genter)
	if genter == ""{
		genter = customer.Genter
	}

	fmt.Printf("年龄(%v)：",customer.Age)
	age := 0
	fmt.Scanln(&age)
	if age == 0 {
		age = customer.Age
	}

	fmt.Printf("电话(%v)：",customer.Phone)
	phone := ""
	fmt.Scanln(&phone)
	if phone == ""{
		phone = customer.Phone
	}
	fmt.Printf("邮箱(%v)：",customer.Email)
	email := ""
	fmt.Scanln(&email)
	if email == ""{
		email = customer.Email
	}
	//修改信息
	if this.customerService.Update(id,name,genter,age,phone,email){
		fmt.Println("----------------修改完成----------------")
	}else{
		fmt.Println("----------------修改失败----------------")
	}

	

}
func (this *customerView) mainMenu(){
	for{
		fmt.Println("----------------客户信息管理软件----------------")
		fmt.Println("                1 添加客户")
		fmt.Println("                2 修改客户")
		fmt.Println("                3 删除客户")
		fmt.Println("                4 客户列表")
		fmt.Println("                5 退出软件")
		fmt.Printf("请选择（1-5）：")
		fmt.Scanln(&this.key)

		switch this.key {
			case "1":
				this.add()
			case "2":
				this.update()
			case "3":
				this.delete()
			case "4":
				this.list()
			case "5":
				this.loop = false
			default:
				fmt.Println("输入错误，请重新输入。。。")
			
		}
		if !this.loop {
			break
		}
	}
}

func main()  {
	fmt.Println("ok")
	customerView := customerView{
		key:"",
		loop:true,
	}
	customerView.customerService = service.NewcustomerService()
	customerView.mainMenu()

	
}