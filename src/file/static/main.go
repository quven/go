package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	file := "./main.go"
	src, err := os.Open(file)
	if err != nil {
		fmt.Println("open file err:", err)
	}
	defer src.Close()

	var letterNum, digitNum, spaceNum int
	reader := bufio.NewReader(src)
	for {
		str, err := reader.ReadString('\n')
		if err != io.EOF {
			for _, s := range str {
				if unicode.IsLetter(s) {
					letterNum++
				}
				if unicode.IsDigit(s) {
					digitNum++
				}
				if unicode.IsSpace(s) {
					spaceNum++
				}
			}
		} else {
			break
		}
	}
	fmt.Printf("统计字母一共有%d个，数字有%d个，空格有%d个\n", letterNum, digitNum, spaceNum)

}
