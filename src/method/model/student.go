package model

type student struct {
	name string
	age  int
}

//student首字母小写，只能用工厂模式在外部访问

func NewStudent(n string, a int) *student {
	return &student{
		name: n,
		age:  a,
	}
}
func NewStudent2() *student {
	return &student{}
}
func (s *student) GetAge() int {
	return s.age
}
func (s *student) SetAge(age int) {
	s.age = age
}
func (s *student) GetName() string {
	return s.name
}
func (s *student) SetName(name string) {
	s.name = name
}
