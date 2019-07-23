package model

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Gender string
}

func NewCustomer(id int, name string, gender string) *Customer {
	return &Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
	}
}
func (c *Customer) Info() string {
	var info string
	info = fmt.Sprintf("%d\t%s\t%s\t\n", c.Id, c.Name, c.Gender)
	return info
}
