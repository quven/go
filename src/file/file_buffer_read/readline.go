package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "test.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file err:", err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	//var line []byte
	for {
		data, prefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		//line = append(line, data...)

		if !prefix {
			fmt.Printf("%s\n", string(data))
		}
	}

}
