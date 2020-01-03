package model

import(
	"fmt"
)

type Customer struct{
	Id int
	Name string
	Genter string
	Age int
	Phone string
	Email string
}

func NewCustomer(Id int, Name string, Genter string, Age int, Phone string, Email string) Customer {
	return Customer{
		Id:Id,
		Name:Name,
		Genter:Genter,
		Age:Age,
		Phone:Phone,
		Email:Email,
	}
}
func NewCustomer2( Name string, Genter string, Age int, Phone string, Email string) Customer {
	return Customer{
		Name:Name,
		Genter:Genter,
		Age:Age,
		Phone:Phone,
		Email:Email,
	}
}
func (this *Customer) GetInfo() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v",this.Id,this.Name,this.Genter,this.Age,this.Phone,this.Email)
}