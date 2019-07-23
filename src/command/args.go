package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行参数一共有", len(os.Args), " 个")

	for i, v := range os.Args {
		fmt.Printf("Args[%d] = %s\n", i, v)
	}
}
