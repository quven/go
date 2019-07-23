package main

import (
	"flag"
	"fmt"
)

func main() {
	var username, password string
	var port int
	flag.StringVar(&username, "u", "xuehao", "用户名，默认为xuehao")
	flag.StringVar(&password, "p", "admin", "密码，默认为admin")
	flag.IntVar(&port, "port", 3306, "端口名，默认为3306")
	flag.Parse()

	fmt.Printf("username=%s, password = %s, port = %d\n", username, password, port)
}
