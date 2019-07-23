package model

import "fmt"

type Account struct {
	Pwd   int
	Price float32
}

//存款方法
func (a *Account) Deposit(pwd int, price float32) {
	if pwd != a.Pwd {
		fmt.Println("您的密码不正确")
		return
	}
	a.Price += price
	fmt.Println("本次您的存款为", price, "元，您的银行卡余额为", a.Price)
}

//取款方法
func (a *Account) Withdraw(pwd int, price float32) {
	if pwd != a.Pwd {
		fmt.Println("您的密码不正确")
		return
	}
	if price > a.Price {
		fmt.Println("余额不足，您的银行卡余额为", a.Price)
		return
	}
	a.Price -= price
	fmt.Println("本次您的取款为", price, "元，您的银行卡余额为", a.Price)
}

//查询银行卡余额方法
func (a *Account) Query(pwd int) {
	if a.Pwd != pwd {
		fmt.Println("密码不正确，请重新输入")
		return
	}
	fmt.Println("您的银行卡余额为", a.Price)
}
