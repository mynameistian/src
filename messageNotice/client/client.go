package main

import (
	_ "bufio"
	"flag"
	"fmt"
	"net"
	_ "os"
	_ "strings"
)

var IP string

func init() {
	flag.StringVar(&IP, "IP", "default", "log in user")
}

func main() {

	flag.Parse() //暂停获取参数
	url := IP + ":8934"
	fmt.Println(url)
	conn, err := net.Dial("tcp", url)
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}

	_, err = conn.Write([]byte("GetBuffSucc"))
	if err != nil {
		fmt.Println("conn.Write err=", err)
	}
	// }

}
