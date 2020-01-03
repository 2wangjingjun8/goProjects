package service

import(
	"go_code/customer/model"
)

type CustomerService struct{
	customers []model.Customer
	customerNum int
}

func NewcustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	Customer := model.NewCustomer(customerService.customerNum,"张三","男",21,"120","1748616485@qq.com")
	customerService.customers = append(customerService.customers,Customer)
	return customerService
}

func (this *CustomerService) List() []model.Customer {
	return this.customers
}
func (this *CustomerService) Add(customer model.Customer) bool {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers,customer)
	return true
}
func (this *CustomerService) Delete(id int) bool {
	index := this.FindByID(id)
	if index == -1 {
		return false
	}
	this.customers = append(this.customers[:index],this.customers[index+1:]...)
	return true
}
func (this *CustomerService) FindByID(id int) int {
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if id == this.customers[i].Id {
			return i
		}
	}
	return index
}
func (this *CustomerService) FindOneByID(id int) model.Customer {
	index := this.FindByID(id)
	if index == -1 {
		return model.Customer{}
	}
	return this.customers[index]
}

func (this *CustomerService) Update(id int,name string,genter string,age int,phone string,email string) bool {
	index := this.FindByID(id)
	if index == -1 {
		return false
	}
	this.customers[index].Name = name
	this.customers[index].Genter = genter
	this.customers[index].Age = age
	this.customers[index].Phone = phone
	this.customers[index].Email = email
	return true
}