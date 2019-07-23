package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "./WriteString.txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("open file err:", err)
		os.Exit(1)
	}
	defer file.Close()
	fileString := "I am Shaun\n"
	file.Seek(0, 2) //内容最后追加
	file.WriteString(fileString)

}
