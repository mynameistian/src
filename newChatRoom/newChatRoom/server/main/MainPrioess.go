package main

import (
	"chatroom.com/message"
	"chatroom.com/server/proess/chatProess"
	"chatroom.com/server/proess/sysProess"
	"chatroom.com/utils/pack"
	"errors"
	"fmt"
	"net"
)

type Proess struct {
	Conn net.Conn
}

func NewProess(conn net.Conn) (proessOb *Proess) {
	proessOb = &Proess{
		Conn: conn,
	}
	return
}

//运行函数
func (proess *Proess) Run() {

	transferOb := pack.NewTransfer(proess.Conn)
	var data []byte
	var inMessageOb message.Message
	var outMessageOb message.Message

	inMessageOb, err := transferOb.ReadPkg()
	if err != nil {
		fmt.Println("transferOb.ReadPkg err is ", err)
		return
	}

	//根据消息类型执行不同功能
	switch inMessageOb.Type {
	case message.LoginMesType:
		//登录
		outMessageOb.Type = message.LoginResMesTYpe
		data, err = sysProess.LoginFun([]byte(inMessageOb.Data), proess.Conn)

	case message.RegisterMesType:
		//注册
		outMessageOb.Type = message.LoginResMesTYpe
		data, err = sysProess.RegisterFun([]byte(inMessageOb.Data))

	case message.RrivateChatMesType:
		//私聊
		outMessageOb.Type = message.RrivateChatMesType
		data, err = chatProess.PrivateChat([]byte(inMessageOb.Data))

	case message.GroupChatMesType:
		//群发
		outMessageOb.Type = message.GroupChatResMesType
		data, err = chatProess.GroupChat([]byte(inMessageOb.Data))

	case message.FriendsListMesType:
		//返回用户列表
		outMessageOb.Type = message.FriendsListResMesType
		data, err = sysProess.FriendsList([]byte(inMessageOb.Data))

	default:
		//错误消息类型
		outMessageOb.Type = message.SysErrResMesType
		err = errors.New("messageType err")
	}

	if err != nil {
		outMessageOb.Code = 404
	} else {
		outMessageOb.Code = 200
	}
	outMessageOb.Data = string(data)
	err = transferOb.WritePkg(outMessageOb)
	if err != nil {
		fmt.Println("transferOb.WritePkg err is ", err)
		return
	}
}
