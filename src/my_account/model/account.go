package model

import (
	"fmt"
)

type Account struct {
	Balance float32
	Note    string
	Record  string
}

func NewAccount(balance float32) *Account {
	return &Account{
		Balance: balance,
		Record:  "类型\t金额\t余额\t说明\n",
	}
}

func (a *Account) Income(money float32, note string) {
	a.Balance += money
	a.Note = note
	a.Record += fmt.Sprintf("收入\t%.2f\t%.2f\t%s\t\n", money, a.Balance, a.Note)
}

func (a *Account) Expenditure(money float32, note string) bool {
	if money > a.Balance {
		fmt.Println("余额不足")
		return false
	}
	a.Balance -= money
	a.Note = note
	a.Record += fmt.Sprintf("支出\t%.2f\t%.2f\t%s\t\n", money, a.Balance, a.Note)
	return true
}
func (a *Account) Show() {
	fmt.Println(a.Record)
}
