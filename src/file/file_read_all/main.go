package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file := "./test.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("read err:", err)
	}
	fmt.Println(string(content))
}
