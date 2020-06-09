package process

import (
	"encoding/json"
	"chatroom.com/common/message"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {

	var loginMes message.LoginMes

	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err :", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.LoginResMesTYpe

	var loginResMes message.LoginResMes

	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 200
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册再使用..."
	}

	data, err = json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal", err)
		return
	}

	resMes.Data = string(data)

	data, err = json.Marshal(resmes)
	if err!= nil {
		fmt.Println("json.Marshal err :" , err)
		return 
	}

	tf:= &util.Transfer{
		Conn : this.Conn
	}

	err = tf.WritePkg(data)
	
	return
}
