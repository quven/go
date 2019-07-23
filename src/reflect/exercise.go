package main

import (
	"fmt"
	"reflect"
)

//简单类型的反射
func reflectTest01(b interface{}) {
	//通过反射来获取传入变量的   type，kind。值
	//1.先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType = ", rType)
	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b)
	n2 := 2 + rVal.Int()

	fmt.Println("n2 = ", n2)
	fmt.Printf("rVal = %v, rVal type=%T\n", rVal, rVal)
	//下面我们将rVal转换成interface
	iv := rVal.Interface()
	//将interface通过断言转换成需要的类型
	num2 := iv.(int)
	fmt.Println("num2 = ", num2)
}

//复杂类型的反射
func reflectTest02(b interface{}) {
	//通过反射获取到传入变量的type  kind值
	rType := reflect.TypeOf(b)
	fmt.Println("rType =", rType)

	//获取到reflectValue
	rVal := reflect.ValueOf(b)

	//获取变量对应的kind
	typeKind := rType.Kind()
	valKind := rVal.Kind()
	fmt.Printf("typeKind = %v, valKind = %v\n", typeKind, valKind)

	//将rVal转换为interface{}
	iv := rVal.Interface()
	fmt.Printf("iv = %v is type =%T\n", iv, iv)
	//把interface{}通过断言转换成需要的类型
	stu, ok := iv.(Student)
	if ok {
		fmt.Println("stu.Name = ", stu.Name)
	}
}
func changeValue(b interface{}) {
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal kind =%v\n", rVal.Kind())
	rVal.Elem().SetInt(1)
}

type Student struct {
	Name string
	Age  int
}

func main() {
	var b int = 10
	reflectTest01(b)
	fmt.Println("-------------")

	stu := Student{
		Name: "tom",
		Age:  12,
	}
	reflectTest02(stu)
	fmt.Println("-------------")
	var c int = 100
	changeValue(&c)
	fmt.Println(c)
}
