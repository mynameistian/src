package main

import (
	"../service"
	"fmt"
)

// CustomerView 界面展示类
type CustomerView struct {
	key       string
	loop      bool
	customers *service.Customers
}

//showView 界面展示
func (customerview *CustomerView) showView() {
	for {
		fmt.Println()
		fmt.Println("---------------------------客户信息管理软件---------------------------------")
		fmt.Println()
		fmt.Println("---------------------------1 添 加 客 户-----------------------------------")
		fmt.Println("---------------------------2 修 改 客 户-----------------------------------")
		fmt.Println("---------------------------3 删 除 客 户-----------------------------------")
		fmt.Println("---------------------------4 客 户 列 表-----------------------------------")
		fmt.Println("---------------------------5 退      出-----------------------------------")
		fmt.Print("请选择(1-5):")
		fmt.Scanln(&customerview.key)

		switch customerview.key {
		case "1":
			customerview.addCustomer()
		case "2":
			customerview.updateCustomer()
		case "3":
			customerview.deleteCustomer()
		case "4":
			fmt.Println()
			fmt.Println("---------------------------- 客 户 列 表-----------------------------------")
			fmt.Println("用户ID\t姓名\t性别\t年龄\t手机号\t电子邮件")
			customerview.customers.ListCustomers()
		case "5":
			customerview.quit()
		default:
			fmt.Println("输入有误，请重新输入")
		}
		customerview.key = "6"
		if customerview.loop == true {
			break
		}
	}
}

//updateCustomer 更改用户信息
func (customerview *CustomerView) updateCustomer() {
	var id int
	var name string
	var gender string
	var age int
	var phone string
	var email string
	for {
		fmt.Println("请输出要更改的客户信息ID[输入-1则退出该更]")
		fmt.Scanln(&id)
		if id == -1 {
			break
		}
		fmt.Println("请输入相关信息:")
		fmt.Print("姓名：")
		fmt.Scanln(&name)
		fmt.Print("性别：")
		fmt.Scanln(&gender)
		fmt.Print("年龄：")
		fmt.Scanln(&age)
		fmt.Print("手机号：")
		fmt.Scanln(&phone)
		fmt.Print("电子邮箱：")
		fmt.Scanln(&email)
		if customerview.customers.UpdateCustomer(id, name, gender, age, phone, email) {
			fmt.Println("更改成功")
		} else {
			fmt.Println("更改失败")
		}
	}
}

// deleteCustomer 删除客户信息
func (customerview *CustomerView) deleteCustomer() {
	var id int
	for {
		fmt.Println("请输出要删除的客户信息ID[输入-1则退出删除]")
		fmt.Scanln(&id)
		if id == -1 {
			break
		}
		if customerview.customers.DeleteCustomer(id) {
			fmt.Println("删除成功")
		} else {
			fmt.Println("删除失败")
		}
	}
}

// addCustomer 添加客户信息
func (customerview *CustomerView) addCustomer() {
	var name string
	var gender string
	var age int
	var phone string
	var email string
	fmt.Println("请输入相关信息:")
	fmt.Print("姓名：")
	fmt.Scanln(&name)
	fmt.Print("性别：")
	fmt.Scanln(&gender)
	fmt.Print("年龄：")
	fmt.Scanln(&age)
	fmt.Print("手机号：")
	fmt.Scanln(&phone)
	fmt.Print("电子邮箱：")
	fmt.Scanln(&email)
	customerview.customers.AddCustomer(name, gender, age, phone, email)
}

// 退出程序
func (customerview *CustomerView) quit() {
	choose := "N"
	fmt.Println("是否要退出程序(Y/N):")
	for {
		fmt.Scanln(&choose)
		if choose == "Y" || choose == "N" || choose == "y" || choose == "n" {
			if choose == "Y" || choose == "y" {
				customerview.loop = true
				break
			}
		} else {
			fmt.Println("输入错误请重新输入[是否要退出程序(Y/N):]:")
		}
	}
	fmt.Println("退出程序")
}

// 主函数
func main() {
	View := CustomerView{}
	View.customers = service.NewCustomers()
	View.showView()
}
