package pack

import (
	"chatroom.com/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type Transfer struct {
	Conn   net.Conn
	Buf    [8096]byte
	BufLen uint32
}

//初始化一个链接
func NewTransfer(conn net.Conn) *Transfer {
	return &Transfer{
		Conn: conn,
	}
}

//接收包
func (triansfer *Transfer) ReadPkg() (mes message.Message, err error) {

	pkgLen, err := triansfer.Conn.Read(triansfer.Buf[:])
	if err != nil {
		return
	}

	fmt.Println("接收包 data is ", triansfer.Buf)

	triansfer.BufLen = binary.BigEndian.Uint32(triansfer.Buf[0:4])
	if uint32(pkgLen) != triansfer.BufLen {
		errString := fmt.Sprintf("buf len err : readbuff[%d] BufLen[%d]\n", triansfer.BufLen, pkgLen)
		err = errors.New(errString)
		return
	}

	err = json.Unmarshal(triansfer.Buf[5:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err :", err)
		return
	}

	return
}

//发送包
func (triansfer *Transfer) WritePkg(mes message.Message) (err error) {

	fmt.Println("发送包 data is ", mes)

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err :", err)
	}

	triansfer.BufLen = uint32(len(data)) + 4
	dataString := fmt.Sprintf("%d%s\n", triansfer.BufLen, data)

	_, err = triansfer.Conn.Write([]byte(dataString))
	if err != nil {
		fmt.Println("this.Conn.Write err :", err)
		return
	}

	return
}
