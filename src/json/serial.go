package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name string
	Age  int
}

func testStruct() {
	var m Monster = Monster{
		Name: "xuehao",
		Age:  123,
	}
	res, err := json.Marshal(m)
	if err != nil {
		fmt.Println("err = ", err)
	}
	fmt.Printf("json = %s\n", res)
}
func testMap() {
	var a map[string]interface{}

	a = make(map[string]interface{})
	a["name"] = "yangchaoyue"
	a["age"] = 13

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("json err is ", err)
	}
	fmt.Printf("map json is %s\n", data)
}
func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}

	m1 = make(map[string]interface{})
	m1["name"] = "name1"
	m1["age"] = 12

	slice = append(slice, m1)
	//fmt.Println(slice)

	var m2 map[string]interface{}

	m2 = make(map[string]interface{})
	m2["name"] = "name2"
	m2["age"] = 12

	slice = append(slice, m2)
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("json err is ", err)
	}
	fmt.Printf("slice json is %s\n", data)
}

func testFloat() {
	var f float32 = 21.21
	data, err := json.Marshal(f)
	if err != nil {
		fmt.Println("json err is ", err)
	}

	fmt.Printf("float json =%s\n", data)
}
func main() {
	testStruct()
	testMap()
	testSlice()
	testFloat()
	return
}
