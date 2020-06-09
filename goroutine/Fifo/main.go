package main

import (
	"fmt"
	_ "runtime"
	_ "sync"
	"time"
)

func inFifo(intChan chan int) {

	for i := 0; i < 50; i++ {
		intChan <- i

		fmt.Printf("输入管道intChan 值为：[%v] \n", i)
	}
	close(intChan)
}

func outFifo(intChan chan int, exitChan chan bool) {

	for {
		res, ok := <-intChan
		if !ok {
			exitChan <- true
			close(exitChan)
			break
		}
		fmt.Printf("读取管道exitChan 值为: [%v]\n", res)
	}
}

func main() {

	var intChan chan int
	intChan = make(chan int, 10)
	var exitChan chan bool
	exitChan = make(chan bool, 1)

	go inFifo(intChan)

	go outFifo(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			fmt.Print("结束\n")
			break
		}
		time.Sleep(time.Second)
	}
}
