package main

import (
	"fmt"
	_ "math/rand"
)

func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("intChan的值%v,intChan的本身地址%v\n", intChan, &intChan)
	//向管道写入数据
	intChan <- 10
	intChan <- 20
	intChan <- 30
	fmt.Printf("len %d,cap %d\n", len(intChan), cap(intChan))
	//var num1, num2 int
	//num1 = <-intChan
	//num2 = <-intChan
	//fmt.Printf("num1=%d, num2=%d\n", num1, num2)
	close(intChan)
	for data := range intChan {
		fmt.Println(data)
	}
	fmt.Printf("len %d,cap %d\n", len(intChan), cap(intChan))
}

type Person struct {
	Name    string
	Age     int
	Address string
}
