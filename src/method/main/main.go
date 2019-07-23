package main

import (
	"fmt"
	"method/model"
)

//判断奇数偶数方法
type int1 int

func (n *int1) CheckNum() {
	if (*n)%2 == 0 {
		fmt.Println("是奇数")
	} else {
		fmt.Println("是偶数")
	}
}

//打印矩形方法
type Util struct {
	Row  int
	Line int
}

func (u *Util) PrintSquare() {
	for i := 0; i < u.Line; i++ {
		for j := 0; j < u.Row; j++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	}
}

//计算器方法
type Calc struct {
	num1 int
	num2 int
	oper string
}

func (c *Calc) Calculation() int {
	var result int
	switch c.oper {
	case "a":
		result = c.num1 + c.num2
		break
	case "b":
		result = c.num1 - c.num2
		break
	case "c":
		result = c.num1 * c.num2
		break
	case "d":
		result = c.num1 / c.num2
		break
	default:
		fmt.Println("计算符号错误")
	}
	return result
}

type Visitor struct {
	Age   int
	Price float32
}

func (v *Visitor) CheckPeople() float32 {
	if (v.Age <= 0) || (v.Age >= 130) {
		panic("年龄不正确")
	}
	if v.Age < 12 {
		v.Price = 0
	} else {
		v.Price = 20
	}
	return v.Price
}
func main() {
	//var n int1 = 12
	//n.CheckNum()

	//fmt.Println("vim-go")
	//打印矩形
	//var s Util = Util{5, 2}
	//s.PrintSquare()

	//一个方法实现加减乘除
	/*
		var c Calc = Calc{11, 2, "d"}
		res := c.Calculation()
		fmt.Println("res = ", res)
	*/
	/*
		var v Visitor
		var age int
		for {
			fmt.Println("请输入年龄")
			fmt.Scanf("%d", &age)
			v.Age = age
			res := v.CheckPeople()
			fmt.Println("门票费用为", res)
		}*/
	/*
			//工厂模式
		var name string = "1234"
		var age int = 11
		var s = model.NewStudent(name, age)

		fmt.Println(s)
		fmt.Println(s.GetAge())
		s.SetAge(100)
		fmt.Println(s)
		fmt.Println(s.GetAge())
		fmt.Println("-----------------")
		var s2 = model.NewStudent2()

		s2.SetName("xuehao")
		s2.SetAge(11)
		fmt.Println("s2 = ", s2)
	*/
	var a model.Account = model.Account{
		Pwd:   123,
		Price: 100,
	}
	a.Withdraw(123, 100)
}
