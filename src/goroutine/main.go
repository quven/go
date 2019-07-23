package main

import (
	"fmt"
	_ "sync"
	"time"
)

/*
var list = make(map[int]int, 10)

var lock sync.Mutex

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
		list[i] = res
	}
}
func main01() {
	for i := 0; i < 200; i++ {
		lock.Lock()
		go test(i)
		lock.Unlock()
	}
	//time.Sleep(10 * time.Second)
	lock.Lock()
	for k, v := range list {
		fmt.Printf("%d:%d\n", k, v)
	}
	lock.Unlock()
}
*/
func go_worker(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println("我的名字是", name)
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "执行完毕")
}
func main() {
	go go_worker("小黑")
	go go_worker("小白")
	for i := 0; i < 5; i++ {
		fmt.Println("我是main")
		time.Sleep(1 * time.Second)
	}
}
