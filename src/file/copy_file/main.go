package main

import (
	_ "bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	srcFile := "./test.txt"
	dstFile := "./xuehao.txt"
	_, err := CopyFile(dstFile, srcFile)
	if err != nil {
		fmt.Println("copy err:", err)
	}
}
func CopyFile(dstFileName, srcFileName string) (int64, error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("open srcfile err:", err)
	}
	defer srcFile.Close()
	//reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open dstfile err:", err)
	}
	//writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	//return io.Copy(writer, reader)
	return io.Copy(dstFile, srcFile)
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func main01() {
	filePath1 := "test.txt"
	filePath2 := "test1.txt"
	data, err := ioutil.ReadFile(filePath1)
	if err != nil {
		fmt.Println("read file err:", err)
	}
	err = ioutil.WriteFile(filePath2, data, 0666)
	if err != nil {
		fmt.Println("writefile err:", err)
	}
	//修改为buffer读写模式

}
