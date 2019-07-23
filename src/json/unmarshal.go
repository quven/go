package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name string `json:"xue_name"`
	Age  int
}

func testStruct() {
	str := "{\"xue_name\": \"xuehao111111\",\"Age\": 17}"
	var m Monster

	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(m)
}

func testMap() {
	str := "{\"name\": \"xuehao_map\",\"age\": 17}"
	var a map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(a)

}
func testSlice() {
	str := "[{\"Name\": \"xuehao1\",\"Age\": 17},{\"Name\": \"xuehao2\",\"Age\": 17}]"
	var s []map[string]interface{}
	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(s)
}
func main() {
	testStruct()
	testMap()
	testSlice()
}
