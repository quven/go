package main

import (
	"encapsulation/model"
	"fmt"
)

func main() {
	var p = model.NewPerson("xuehao")
	p.SetAge(11)
	fmt.Println(p)
}
