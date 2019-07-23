package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (m *Monster) Store() bool {
	jsonRes, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json Marshal error:", err)
		os.Exit(1)
	}
	fileInput := fmt.Sprintf("%s", jsonRes)
	//文件写入
	fileName := "test.txt"
	res := m.WriteFile(fileName, fileInput)
	return res
}

//写入文件
func (m *Monster) WriteFile(fileName, content string) bool {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("open file error:", err)
		return false
	}
	defer file.Close()
	fileWrite := bufio.NewWriter(file)
	fileWrite.WriteString(content)
	fileWrite.Flush()
	return true
}
func (m *Monster) ReStore() Monster {
	fileName := "test.txt"
	content := m.GetFileContent(fileName)
	var result Monster
	err := json.Unmarshal([]byte(content), &result)
	if err != nil {
		fmt.Println("json unmarshal err:", err)
	}
	return result
}

//获取文件内容
func (m *Monster) GetFileContent(fileName string) string {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("open file err:", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var result string
	for {
		str, err := reader.ReadString('\r')
		result += str
		if err == io.EOF {
			return result
		}
	}

}
func main() {
	var m Monster = Monster{
		Name:  "test",
		Age:   12,
		Skill: "fire",
	}
	m.ReStore()
}
