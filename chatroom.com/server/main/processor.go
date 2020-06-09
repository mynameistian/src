package main

import (
	"chatroom.com/common/message"
	"chatroom.com/server/process"
	"chatroom.com/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) serverProcessMes(conn net.Conn, mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		up := &process2.User
		err = serverProcessMes(conn, mes)
	case message.RegisterMesType:
		//处理注册
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}

	return
}
