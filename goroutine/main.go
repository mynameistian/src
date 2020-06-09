package main

import (
	"fmt"
	"time"
)

func goroutineFunc() {
	var sBuf string = "123"
	go updateString(&sBuf)
	//time.Sleep(time.Second * 3)
	sBuf = "789"
	time.Sleep(time.Second * 1)
	fmt.Printf(" goroutineFunc sBuf is [%v] sBuf 的 地址为[%p]\n ", sBuf, &sBuf)
	return
}

func updateString(sBuf *string) {
	*sBuf = "345"
	time.Sleep(time.Second * 3)
	fmt.Printf(" updateString sBuf is [%v] sBuf 的 地址为[%p] \n ", *sBuf, sBuf)
	return
}

func main() {
	go goroutineFunc()
	time.Sleep(time.Second * 20)
	fmt.Println("结束！")
	return
}
