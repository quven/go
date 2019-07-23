package main

import (
	"fmt"
	"my_account/model"
)

func main() {
	var balance float32
	var note string
	var money float32
	var loop bool = false
	balance = 10000
	account := model.NewAccount(balance)
	//account.Income(100, "工资")
	//account.Show()
	var key int

	for {
		fmt.Println("--------家庭收支记账软件--------")
		fmt.Println("        1:收支明细")
		fmt.Println("        2:登记收入")
		fmt.Println("        3:登记支出")
		fmt.Println("        4:退出软件")
		fmt.Printf("请选择你的操作:")
		fmt.Scanln(&key)
		switch key {
		case 1:
			account.Show()
			break
		case 2:
			fmt.Printf("请输入收入金额:")
			fmt.Scanln(&money)
			fmt.Printf("请输入收入说明:")
			fmt.Scanln(&note)
			account.Income(money, note)
			break
		case 3:
			fmt.Printf("请输入支出金额:")
			fmt.Scanln(&money)
			fmt.Printf("请输入支出说明:")
			fmt.Scanln(&note)
			res := account.Expenditure(money, note)
			if !res {
				break
			}
			break
		case 4:
			loop = true
			break
		default:
			fmt.Printf("请输入正确的数字")
		}
		if loop == true {
			break
		}
	}
}
