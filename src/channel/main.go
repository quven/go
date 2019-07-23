package main

import (
	"fmt"
	_ "time"
)

func WriteData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Println("写入数据", i)
		//time.Sleep(time.Second)
	}
	close(intChan)
}
func ReadData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("读取数据", v)
		//time.Sleep(time.Second)
	}
	//读取完数据后，即任务完成
	exitChan <- true
	close(exitChan)
}
func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go WriteData(intChan)
	go ReadData(intChan, exitChan)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
