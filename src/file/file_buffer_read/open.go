package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "./test.txt"
	file, err := os.Open(fileName)
	//file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("open file err:", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			//文件读取结束
			//fmt.Println("read err:", err)
			return
		}
		fmt.Printf("%s", str)
	}
}
