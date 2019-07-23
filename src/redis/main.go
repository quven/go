package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis" //引入redis包
)

func main() {
	con, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis err:", err)
		return
	}
	defer con.Close()
	//写入内容
	_, err = con.Do("set", "name", "test")
	if err != nil {
		fmt.Println("set error:", err)
	}
	fmt.Println("连接redis成功:", con)
}
