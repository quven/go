package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		//等待客户端通过conn发送信息
		//如果客户端没有Write[发送]，那么协程就阻塞在这里
		//fmt.Printf("服务器在等待客户端%s, 发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("客户端退出")
			return
		}
		if err != nil {
			fmt.Println("服务器Read err", err)
			return
		}
		//显示客户端发送的内容到服务器的终端
		fmt.Println(string(buf[:n]))
	}
}

func Server() {
	//Listen函数创建的服务器
	//头层皮：网络协议
	//127.0.0.1:8888    8888本机ip和端口
	l, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer l.Close()
	//循环等待客户端访问
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("accept error", err)
		}
		fmt.Printf("访问客户端信息：con = %v 客户端ip=%v\n", conn, conn.RemoteAddr())
		go process(conn)
	}
}

func main() {
	Server()
}
