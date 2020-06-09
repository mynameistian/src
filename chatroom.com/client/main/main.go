package main

import (
	"fmt"
	"os"
)

var userId int
var userPwd string

func main() {

	var key int
	var loop = true
	for loop {
		fmt.Println("------------------------------------------------------------------")
		fmt.Println("\t\t 1登录聊天室")
		fmt.Println("\t\t 2注册用户")
		fmt.Println("\t\t 3退出系统")
		fmt.Println("\t\t 请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("您的输入有误，请重新输入")

		}
	}

	if key == 1 {
		fmt.Println("请输入用户的id")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户的密码")
		fmt.Scanf("%s\n", &userPwd)\
		err := login(userId, userPwd)
		if err != nil {
			fmt.Println("登录失败")
		}else{
			fmt.Println("登录成功")
		}
	}else if key == 2{
		fmt.Println("进行用户注册的逻辑...")
	}
}