package service

import (
	"../model"
	"fmt"
)

//Customers 结果体
type Customers struct {
	customers   []model.Customer
	customerNum int
}

// NewCustomers 初始化
func NewCustomers() *Customers {
	Customers := &Customers{}
	Customers.customerNum = 0
	customer := model.NewCustomer(0, " ", " ", 0, " ", " ")
	Customers.customers = append(Customers.customers, customer)
	return Customers
}

//AddCustomer 添加客户信息
func (customers *Customers) AddCustomer(name string, gender string, age int, phone string, email string) bool {
	customers.customerNum++
	customer := model.NewCustomer(customers.customerNum,
		name,
		gender,
		age,
		phone,
		email,
	)
	customers.customers = append(customers.customers, customer)
	return true
}

//UpdateCustomer 更改用户信息
func (customers *Customers) UpdateCustomer(id int, name string, gender string, age int, phone string, email string) bool {
	index := customers.FindIndex(id)
	if index == -1 {
		return false
	}
	customers.customers[index].Name = name
	customers.customers[index].Gender = gender
	customers.customers[index].Age = age
	customers.customers[index].Phone = phone
	customers.customers[index].Email = email
	return true
}

//FindIndex 根据ID 查找索引
func (customers *Customers) FindIndex(ID int) (index int) {
	for i, customer := range customers.customers {
		if customer.ID == ID {
			index = i
			return
		}
	}
	index = -1
	return
}

//DeleteCustomer 删除用户
func (customers *Customers) DeleteCustomer(ID int) bool {
	index := customers.FindIndex(ID)
	if index == -1 {
		return false
	}
	customers.customers = append(customers.customers[:index], customers.customers[index+1:]...)
	return true
}

// ListCustomers 展示客户信息
func (customers *Customers) ListCustomers() {
	for _, customer := range customers.customers {
		if customer.ID == 0 {
			continue
		}
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t\t%v \n", customer.ID, customer.Name, customer.Gender, customer.Age, customer.Phone, customer.Email)
	}
}
