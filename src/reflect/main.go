package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("-----start-------")
	fmt.Println(s)
	fmt.Println("-----end-------")
}
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}) {
	rType := reflect.TypeOf(a)
	rVal := reflect.ValueOf(a)
	kd := rVal.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	num := rVal.NumField()
	fmt.Println("struct fields:", num)
	//变量结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("field %d:值为:%v\n", i, rVal.Field(i))
		tagVal := rType.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("field %d: tag为=%v\n", i, tagVal)
		}
	}

	//结构体的方法操作
	numOfMethod := rVal.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	//var params []reflect.Value
	rVal.Method(1).Call(nil)

	//调用结构体的第1个方法Method(0)
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := rVal.Method(0).Call(params)
	fmt.Println("res = ", res[0].Int())
	//反射修改字段

}
func main() {
	var a Monster = Monster{
		Name:  "tom",
		Age:   100,
		Score: 30.1,
		Sex:   "男",
	}
	TestStruct(a)
}
