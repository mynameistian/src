package model

// Customer 结构体
type Customer struct {
	ID     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

// NewCustomer 初始化函数
func NewCustomer(id int, name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		ID:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}
