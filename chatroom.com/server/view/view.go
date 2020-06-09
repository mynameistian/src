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

	fmt.Printf("\n\n----------------------------------欢迎进入聊天室管理后台---------------------------------------\n\n")
	fmt.Printf("                                 1: 管理员登录\n")
	fmt.Printf("                                 2: 管理员注册\n")
	fmt.Printf("                                 3: 退出程序\n")
	fmt.Printf("请选择(1-3):\n")
	fmt.Scanln(&menu.Key)
}

//LoginView 登录界面
func (menu *Menu) LoginView() {
	fmt.Printf("------------------------------------管理员登录------------------------------------------\n\n")
	fmt.Printf("                                 1:请输入账号与密码\n")
	fmt.Printf("账号:")
	fmt.Scanln(&menu.Useid)
	fmt.Printf("密码")
	fmt.Scanln(&menu.Pwd)
	return
}

//MangeView 管理界面
func (menu *Menu) MangeView() {
	fmt.Printf("------------------------------------程序管理界面------------------------------------------\n\n")
	fmt.Printf("                                 1: 开启服务\n")
	fmt.Printf("                                 2: 暂停服务\n")
	fmt.Printf("                                 3: 关闭服务\n")
	fmt.Printf("                                 4: 返回上一层\n")
	fmt.Printf("请选择(1-4):\n")
	fmt.Scanln(&menu.Key)
	return
}

//RegisterView 注册界面
func (menu *Menu) RegisterView() {
	fmt.Printf("------------------------------------管理员注册------------------------------------------\n\n")
	fmt.Printf("                                 1:请输入账号、名称、密钥与密码\n")
	fmt.Printf("账号:")
	fmt.Scanln(&menu.Useid)
	fmt.Printf("姓名：")
	fmt.Scanln(&menu.Key)
	fmt.Printf("密码：")
	fmt.Scanln(&menu.Pwd)
	return
}

//Newmenu 初始化
func Newmenu() *Menu {
	menu := &Menu{
		Key:      0,
		Useid:    "000000",
		Name:     "",
		Pwd:      "000000",
		QuitType: false,
	}
	return menu
}
