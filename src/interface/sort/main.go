package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Student struct {
	Name string
	Age  int
}
type StudentSlice []Student

func (s StudentSlice) Len() int {
	return len(s)
}
func (s StudentSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}
func (s StudentSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func main() {
	var result StudentSlice
	for i := 0; i < 8; i++ {
		tmp := Student{
			Name: fmt.Sprintf("Name%d", rand.Intn(10)),
			Age:  rand.Intn(100),
		}
		result = append(result, tmp)
	}
	//排序前
	for _, v := range result {
		fmt.Printf("姓名：%s, 年龄：%d\n", v.Name, v.Age)
		//fmt.Println(v)
	}
	fmt.Println("排序后")
	sort.Sort(result)
	for _, v := range result {
		fmt.Printf("姓名：%s, 年龄：%d\n", v.Name, v.Age)
		//fmt.Println(v)
	}
}
