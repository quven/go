package service

import (
	"CustomerManager/model"
	_ "fmt"
)

type CustomerService struct {
	List []model.Customer
}

//初始化CustomerService对象
func NewCustomerService() *CustomerService {
	customer_service := &CustomerService{}
	//customer := model.NewCustomer(1, "xuehao", "man")
	//customer_service.List = append(customer_service.List, *customer)
	return customer_service
}

//获取客户列表
func (c *CustomerService) GetList() []model.Customer {
	return c.List
}

//添加客户
func (c *CustomerService) Add(id int, name string, gender string) {
	customer := model.NewCustomer(id, name, gender)
	c.List = append(c.List, *customer)
}

//查找某位用户
func (c *CustomerService) Find(id int) model.Customer {
	list := c.GetList()
	for i := 0; i < len(list); i++ {
		if list[i].Id == id {
			return list[i]
		}
	}
	return list[len(list)]
}

//修改用户属性
func (c *CustomerService) Edit(id int, name, gender string) {
	list := c.GetList()
	for i := 0; i < len(list); i++ {
		if id == list[i].Id {
			list[i].Name = name
			list[i].Gender = gender
		}
	}
}

//删除用户
func (c *CustomerService) Del(id int) {
	list := c.GetList()
	for i := 0; i < len(list); i++ {
		if list[i].Id == id {
			//删除切片元素
			j := i + 1
			c.List = append(c.List[:i], c.List[j:]...)
		}
	}
}
