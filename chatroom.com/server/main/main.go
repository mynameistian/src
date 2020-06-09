package main

import (
	"chatroom.com/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
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

	_, err = json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal", err)
		return
	}
	return
}

func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		err = serverProcessMes(conn, mes)
	case message.RegisterMesType:
		//处理注册
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}

	return
}

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据 ")

	_, err = conn.Read(buf[:4])
	if err != nil {
		return
	}

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}

	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err :", err)
		return
	}

	return

}
func process(conn net.Conn) {
	defer conn.Close()

	for {

		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也退出")
				return
			} else {
				fmt.Println("readPkg err :", err)
				return
			}
		}
		err = serverProcessMes(conn, &mes)
		if err != nil {
			return
		}
	}
}

func main() {
	//提示信息
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err :", err)
		return
	}

	for {
		fmt.Println("等待客户端来链接服务器...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err:", err)
		}

		go process(conn)
	}
}
