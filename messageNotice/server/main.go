package main

import (
	"fmt"
	"io/ioutil"
	"net" //做网络socket开发时,net包含有我们需要所有的方法和函数
)

var num int

//自己编写一个函数，接收两个文件路径 srcFileName dstFileName
func CopyFile(srcFileName string, dstFileName string) {

	//将d:/abc.txt 文件内容导入到  e:/kkk.txt
	//1. 首先将  d:/abc.txt 内容读取到内存
	//2. 将读取到的内容写入 e:/kkk.txt
	fmt.Println(srcFileName)
	fmt.Println(dstFileName)
	data, err := ioutil.ReadFile(srcFileName)
	if err != nil {
		//说明读取文件有错误
		fmt.Printf("read file err=%v\n", err)
		return
	}
	err = ioutil.WriteFile(dstFileName, data, 0666)
	if err != nil {
		fmt.Printf("write file error=%v\n", err)
	}
}

func process(conn net.Conn) {

	//这里我们循环的接收客户端发送的数据
	defer conn.Close() //关闭conn
	srcFileName := "C:\\Program Files\\CASIC\\HDmon Agent\\HDmon_agentd.log"
	dstFileName := "C:\\Program Files\\CASIC\\HDmon Agent\\bak\\HDmon_agentd.log"

	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		n, err := conn.Read(buf) //从conn读取
		if err != nil {
			fmt.Printf("客户端退出 err=%v", err)
			return //!!!
		}
		//3. 显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
		if string(buf[:n]) == "GetBuffSucc" {
			fmt.Println("CopyFile")
			dstFileName1 := fmt.Sprintf("%s-%d", dstFileName, num)
			CopyFile(srcFileName, dstFileName1)
			num++
		}
	}
}

func main() {

	fmt.Println("服务器开始监听....")
	//net.Listen("tcp", "0.0.0.0:8888")
	//1. tcp 表示使用网络协议是tcp
	//2. 0.0.0.0:8888 表示在本地监听 8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8934")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close() //延时关闭listen

	//循环等待客户端来链接我
	for {
		//等待客户端链接
		fmt.Println("等待客户端来链接....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)

		} else {
			fmt.Printf("Accept() suc con=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		//这里准备其一个协程，为客户端服务
		go process(conn)
	}
}
