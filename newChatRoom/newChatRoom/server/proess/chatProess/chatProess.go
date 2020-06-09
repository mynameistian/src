package chatProess

import (
	"chatroom.com/message"
	"chatroom.com/server/proess/sysProess"
	"chatroom.com/utils/pack"
	"encoding/json"
	"fmt"
)

//私聊函数
func PrivateChat(data []byte) (Outdata []byte, err error) {
	var messageOb message.MessageData

	err = json.Unmarshal(data, &messageOb)
	if err != nil {
		fmt.Println("json.Unmarshal err is ", err)
		return
	}

	//获取接收者的套接字
	OnlineList := sysProess.OnlineList.FindUserInfo(messageOb.ReceiverId)

	TransferOb := pack.NewTransfer(OnlineList.Conn)

	var Outmessage message.Message
	Outmessage.Type = message.RrivateChatMesType
	Outmessage.Data = string(Outdata)

	err = TransferOb.WritePkg(Outmessage)
	if err != nil {
		fmt.Println("TransferOb.WritePkg err is ", err)
		return
	}

	return
}

//群发函数
func GroupChat(data []byte) (Outdata []byte, err error) {

	var messageOb message.MessageData

	err = json.Unmarshal(data, &messageOb)
	if err != nil {
		fmt.Println("json.Unmarshal err is ", err)
		return
	}

	//获取接收者的套接字
	OnlineList := sysProess.OnlineList.FriendsList()
	for _, v := range OnlineList {

		TransferOb := pack.NewTransfer(v.Conn)
		var Outmessage message.Message
		Outmessage.Type = message.RrivateChatMesType
		Outmessage.Data = string(Outdata)
		err = TransferOb.WritePkg(Outmessage)
		if err != nil {
			fmt.Println("TransferOb.WritePkg err is ", err)
			continue
		}
	}
	return
}
