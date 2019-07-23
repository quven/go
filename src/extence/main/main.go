package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

func (s *Student) ShowInfo() {
	fmt.Printf("姓名：%s. 年龄:%d, 分数：%d\n", s.Name, s.Age, s.Score)
}

type Pupil struct {
	Student
}

func (p *Pupil) Testing() {
	fmt.Println(p.Name, "is testing")
}
func main() {
	/*
		var p Pupil = Pupil{
			Name:  "xuehao",
			Age:   11,
			Score: 98,
		}*/

	var p Pupil
	p.Name = "xuehao"
	p.Age = 11
	p.Score = 98
	p.Testing()
	p.ShowInfo()
}
