package main

import (
	"chatroom.com/server/model"
	"chatroom.com/server/proess/sysProess"
	"fmt"
	"net"
	"time"
)

//初始化
func init() {
	model.InitRedis(16, 0, 300*time.Second, "localhost:6379") //数据库初始化
	sysProess.InitOnlineList()                                //全局变量初始化
}

//主函数
func main() {
	//开启主进程
	MainProess()
}

func MainProess() {
	//进入主进程
	fmt.Println("进入主进程.....")
	//绑定ip 端口
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err is ", err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Acceptis ", err)
			continue
		}
		proessOb := NewProess(conn)
		go proessOb.Run()
	}

}
