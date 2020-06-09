package sysProess

import (
	"chatroom.com/message"
	"chatroom.com/utils/pack"
	"encoding/json"
	"fmt"
	"net"
)

type sysProess struct {
	Conn net.Conn
}

func NewSysProess(inConn net.Conn) *sysProess {
	sysproess := &sysProess{
		Conn: inConn,
	}

	return sysproess
}

func (sysproess *sysProess) Login(userId int, userPwd string) (UserName string, err error) {

	loginMes := message.LoginMes{
		UserId:  userId,
		UserPwd: userPwd,
	}
	data, err := json.Marshal(loginMes)
	if err != nil {
		return
	}
	messageData := message.Message{
		Type: message.LoginMesType,
		Data: string(data),
		Code: 000,
	}

	Buf, err := json.Marshal(messageData)
	if err != nil {
		return
	}

	var packOb pack.Transfer
	packOb.Conn = sysproess.Conn
	err = packOb.WritePkg(Buf)
	if err != nil {
		return
	}

	var readMess message.Message
	readMess, err = packOb.ReadPkg()
	if err != nil {
		return
	}

	if readMess.Code == 200 {
		var logResMess message.LoginResMes
		err = json.Unmarshal([]byte(readMess.Data), &logResMess)
		if err != nil {
			fmt.Println("json.Unmarshal err is :", err)
			return
		}
		UserName = logResMess.UserName
	}

	return

}

func (sysproess *sysProess) registered(inConn net.Conn) (err error) {

}

func (sysproess *sysProess) friendsList(inConn net.Conn) (err error) {

}
