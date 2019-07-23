package main

import "fmt"

func main() {
	var key int
	var loop bool = true

	//本金
	balance := 10000
	money := 0
	note := ""
	record := "类型\t余额\t金额\t说明\n"
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
			fmt.Printf(record)
			break
		case 2:
			fmt.Printf("请输入收入金额:")
			fmt.Scanln(&money)
			fmt.Printf("请输入收入说明:")
			fmt.Scanln(&note)
			balance += money
			record = record + fmt.Sprintf("收入\t%d\t%d\t%s\n", balance, money, note)
			break
		case 3:
			fmt.Printf("请输入支出金额:")
			fmt.Scanln(&money)
			fmt.Printf("请输入支出说明:")
			fmt.Scanln(&note)
			if money > balance {
				fmt.Println("余额不足")
				break
			}
			balance -= money
			record = record + fmt.Sprintf("支出\t%d\t%d\t%s\n", balance, money, note)
			break
		case 4:
			loop = false
			fmt.Println("您已经退出该软件")
			break
		default:
			fmt.Println("请输入正确的选项")
			break
		}
		if loop == false {
			break
		}
	}
}
