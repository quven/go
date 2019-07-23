package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	con, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("connect err", err)
		return
	}
	fmt.Println("con 成功=", con)

	for {
		//客户端可以发送单行数据，然后就退出
		reader := bufio.NewReader(os.Stdin) //os.Stdin代表标准输入[终端]

		//从终端读取一行用户输入，并准备发送给服务端
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readstring err = ", err)
		}
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			break
		}
		//再将line发送给服务器
		n, err := con.Write([]byte(line))
		if err != nil {
			fmt.Println("con.write err = ", err)
		}
		fmt.Printf("发送了%d字节到服务器\n", n)
	}
	fmt.Println("程序退出")
}
