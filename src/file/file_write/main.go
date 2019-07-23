package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	filePath := "./test.txt"
	//file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	//file, err := os.OpenFile(filePath, os.O_TRUNC|os.O_WRONLY, 0666)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("err is :", err)
		os.Exit(1)
	}
	defer file.Close()
	//读取原来的文件
	reader := bufio.NewReader(file)
	for {
		res, err := reader.ReadString('\r')
		if err == io.EOF {
			break
		}
		fmt.Println(res)
	}

	str := "北京早安 hello\r\n"
	writer := bufio.NewWriter(file)
	//写入时，使用带缓存的buffer
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//writer是带缓存，因此在调用WrteString方法时，其实内容是先写入缓存的，所以需要调用Flush方法，将缓存中的数据真正的写入到文件中，否则文件中没有数据
	writer.Flush()

}
