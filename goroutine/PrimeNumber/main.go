package main

import (
	"fmt"
	"time"
)

func ifSuShu(num int) bool {
	for v := 2; v < num-1; v++ {
		if num%v == 0 {
			return false
		}
	}
	return true
}

func intNumFifo(numChan chan int) {

	for v := 0; v < 800000; v++ {
		numChan <- v
		fmt.Printf("输入管道intChan 值为：[%v] \n", v)
	}
	close(numChan)
}

func primeNumFifo(numChan chan int, primeNumChan chan int, resultChan chan int, tid int) {
	for v := range numChan {
		if ifSuShu(v) {
			fmt.Printf("素数为[%v]\n", v)
			primeNumChan <- v
		}
	}
	resultChan <- tid
}

func putPrimeNumChan(primeNumChan chan int, resultChan chan int) {

	var Fifo0 bool
	var Fifo1 bool
	var Fifo2 bool
	var Fifo3 bool
	for {
		i, ok := <-resultChan
		if !ok {
			time.Sleep(time.Second)
			continue
		}
		switch i {
		case 0:
			Fifo0 = true
		case 1:
			Fifo1 = true
		case 2:
			Fifo2 = true
		case 3:
			Fifo3 = true
		}
		if Fifo0 && Fifo1 && Fifo2 && Fifo3 {
			close(primeNumChan)
			close(resultChan)
			fmt.Println("结束！")
			break
		}
	}
	h := 0
	num := 1
	fmt.Println("素数结果:")
	for v := range primeNumChan {
		h++
		num++
		if h == 20 {
			fmt.Println()
			h = 0
		}
		fmt.Printf("[%v]\t", v)
	}
	fmt.Printf("\n素数个数为[%v]", num)
}

func main() {
	numChan := make(chan int, 2000000)
	primeNumChan := make(chan int, 2000000)
	resultChan := make(chan int, 4)

	go intNumFifo(numChan)
	for tid := 0; tid <= 3; tid++ {
		go primeNumFifo(numChan, primeNumChan, resultChan, tid)
	}
	putPrimeNumChan(primeNumChan, resultChan)
}
