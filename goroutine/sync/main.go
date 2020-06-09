package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)
	lock  sync.Mutex
)

func test() {
	for i := 0; i <= 10; i++ {
		fmt.Println("test() Hello test !")
		time.Sleep(time.Second)
	}
}

func test02(n int) {
	num := 1
	for i := 1; i <= n; i++ {
		num *= i
	}
	lock.Lock()
	myMap[n] = num
	lock.Unlock()
}
func main() {
	cpuNum := runtime.NumCPU()
	fmt.Printf("cpuNum ：[%v]", cpuNum)
	for i := 0; i <= 20; i++ {
		go test02(i)
	}
	time.Sleep(time.Second * 5)
	lock.Lock()

	for i, v := range myMap {
		fmt.Printf("myMap[%d] is [%v]\n", i, v)
	}
	lock.Unlock()
}
