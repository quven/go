package main

import "fmt"

/*
type Usb interface {
	insert()
	update()
}
type Phone struct {
}

func (p *Phone) insert() {
	fmt.Println("pohoe is inserting")
}
func (p *Phone) update() {
	fmt.Println("pohoe is updating")
}

type Camera struct {
}

func (c *Camera) insert() {
	fmt.Println("camera is inserting")
}
func (p *Camera) update() {
	fmt.Println("camera is updating")
}

type Computer struct {
}

func (c *Computer) working(u Usb) {
	u.insert()
	u.update()
}
func main() {
	//var p Phone
	//var c Computer
	c := &Computer{}
	p := &Phone{}
	c.working(p)
}
*/
type Humaner interface {
	SayHi()
}

type Student struct {
	Name string
}

func (s *Student) SayHi() {
	fmt.Printf("student %s is say hi\n", s.Name)
}

type Teacher struct {
	Name string
}

func (t *Teacher) SayHi() {
	fmt.Printf("teacher %s is say hi\n", t.Name)
}

func WhoSayHi(i Humaner) {
	i.SayHi()
}

type MyStr string

func (t MyStr) SayHi() {
	fmt.Printf("string %s is say hi\n", t)
}
func main() {
	//普通调用方法
	var s *Student = &Student{
		Name: "xiaoming",
	}
	s.SayHi()
	var t *Teacher = &Teacher{
		Name: "hanmei",
	}
	t.SayHi()
	var str MyStr = "1111111"
	str.SayHi()
	//接口形式调用方法
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(str)
	fmt.Println("-------------")
	x := make([]Humaner, 3)
	x[0], x[1], x[2] = s, t, str
	for _, value := range x {
		value.SayHi()
	}

}
