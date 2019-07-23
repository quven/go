package main

import (
	"CustomerManager/service"
	"fmt"
)

type View struct {
	Key             string
	Loop            bool
	CustomerService *service.CustomerService
}

func NewView() *View {
	return &View{Loop: false}
}

//客户列表
func (v *View) List() {
	customer_list := v.CustomerService.GetList()
	fmt.Println("编号\t姓名\t性别")
	var info string
	for i := 0; i < len(customer_list); i++ {
		info = customer_list[i].Info()
		fmt.Println(info)
	}
	fmt.Println("用户列表结束\n")
}

//添加客户
func (v *View) Add() {
	var id int
	var name, gender string
	fmt.Printf("请输入用户id:")
	fmt.Scanln(&id)
	fmt.Printf("请输入用户姓名:")
	fmt.Scanln(&name)
	fmt.Printf("请输入用户性别:")
	fmt.Scanln(&gender)
	fmt.Println(id, name, gender)
	v.CustomerService.Add(id, name, gender)

}

//修改客户
func (v *View) Edit() {
	var id int
	var name, gender string
	fmt.Printf("请输入需要修改的用户id:")
	fmt.Scanln(&id)

	customer := v.CustomerService.Find(id)
	fmt.Printf("请输入需要修改的姓名,原名为(%s)", customer.Name)
	fmt.Scanln(&name)
	fmt.Printf("请输入需要修改的性别,原来为(%s)", customer.Gender)
	fmt.Scanln(&gender)
	fmt.Println(name, gender)
	v.CustomerService.Edit(id, name, gender)
}

//删除客户
func (v *View) Del() {
	var id int
	fmt.Printf("请输入需要删除的用户的id:")
	fmt.Scanln(&id)
	v.CustomerService.Del(id)

}
func (v *View) ShowMenu() {
	//初始化数据
	v.CustomerService = service.NewCustomerService()
	for {
		fmt.Println("--------客户关系管理系统---------------")
		fmt.Println("           1:查看客户列表")
		fmt.Println("           2:增加客户")
		fmt.Println("           3:修改客户")
		fmt.Println("           4:删除客户")
		fmt.Println("           5:退出系统")
		fmt.Printf("请输入您想要的操作:")
		fmt.Scanln(&v.Key)
		switch v.Key {
		case "1":
			v.List()
			break
		case "2":
			v.Add()
			break
		case "3":
			v.Edit()
			break
		case "4":
			v.Del()
			break
		case "5":
			fmt.Printf("您确定要退出系统嘛？y\\n\n")
			fmt.Scanln(&v.Key)
			if v.Key == "y" {
				v.Loop = true
			}
			break
		default:
			fmt.Println("输入错误，请重新输入")
		}
		if v.Loop == true {
			fmt.Println("您已经退出客户管理系统")
			break
		}
	}
}
func main() {
	view := NewView()

	view.ShowMenu()
}
