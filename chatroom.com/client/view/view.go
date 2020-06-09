package view

import (
	"fmt"
)

//Menu 菜单类
type Menu struct {
	Key      int
	Useid    string
	Name     string
	Pwd      string
	QuitType bool
}

//MenuView 主菜单
func (menu *Menu) MenuView() {

	fmt.Printf("\n\n----------------------------------欢迎进入聊天室---------------------------------------\n\n")
	fmt.Printf("                                 1: 登录账号\n")
	fmt.Printf("                                 2: 注册账号\n")
	fmt.Printf("                                 3: 退出程序\n")
	fmt.Printf("请选择(1-3):\n")
	fmt.Scanln(&menu.Key)
}

//LoginView 登录界面
func (menu *Menu) LoginView() {
	fmt.Printf("\n------------------------------------登录账号------------------------------------------\n\n")
	fmt.Printf("                                 1:请输入账号与密码\n")
	fmt.Printf("账号:")
	fmt.Scanln(&menu.Useid)
	fmt.Printf("密码")
	fmt.Scanln(&menu.Pwd)
	return
}

//MangeView 管理界面
func (menu *Menu) MangeView() {
	fmt.Printf("------------------------------------登录账号------------------------------------------\n\n")
	fmt.Printf("-------------------------------------------------------------------------------------\n")
	fmt.Printf("在线好友:\n")
}

//RegisterView 注册界面
func (menu *Menu) RegisterView() {
	fmt.Printf("------------------------------------注册账号------------------------------------------\n\n")
	fmt.Printf("                                 1:请输入账号、姓名与密码\n")
	fmt.Printf("账号:")
	fmt.Scanln(&menu.Useid)
	fmt.Printf("姓名")
	fmt.Scanln(&menu.Name)
	fmt.Printf("密码")
	fmt.Scanln(&menu.Pwd)
	return
}

//NewMenu 初始化
func NewMenu() *Menu {
	menu := &Menu{
		Key:      0,
		Useid:    "000000",
		Name:     "",
		Pwd:      "000000",
		QuitType: false,
	}
	return menu
}
