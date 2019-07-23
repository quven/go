package model

//import "fmt"

type person struct {
	Name string
	age  int
	sal  float64
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	p.age = age
}
func (p *person) GetAge() int {
	return p.age
}
