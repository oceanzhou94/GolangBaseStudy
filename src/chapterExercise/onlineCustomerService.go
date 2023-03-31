/*编程实现模拟在线客服*/
package main

import (
	"fmt"
)

var balance float64 = 0.00 //话费余额
var flow float64 = 0.00    // 流量
var shortMessage int = 0   //短信

//功能菜单
func mainMenu() {
	fmt.Println("业务查询请按1")
	fmt.Println("手机充值请按2")
	fmt.Println("业务办理请按3")
	fmt.Println("语音导航请按4")
	fmt.Println("人工服务请按5")
	fmt.Println("退出服务请按0")
}

//业务查询二级菜单
func queryMenu() {
	fmt.Println("业务查询二级菜单")
	fmt.Println("流量查询请按1")
	fmt.Println("话费查询请按2")
	fmt.Println("短信查询请按3")
	fmt.Println("返回上级请按0")
}

//手机充值二级菜单
func chargeMenu() {
	fmt.Println("手机充值二级菜单")
	fmt.Println("流量充值请按1")
	fmt.Println("话费充值请按2")
	fmt.Println("短信充值请按3")
	fmt.Println("返回上级请按0")
}

//业务查询逻辑函数
func businessQuery() {
	var userChoice int = 1
	for userChoice != 0 {
		queryMenu()
		fmt.Printf("请输入您的选择：")

		_, inputError := fmt.Scanln(&userChoice)
		//对用户输入校验
		if inputError != nil {
			fmt.Printf("输入错误：%v\n", inputError)
		} else {
			switch userChoice {
			case 1:
				//1.流量查询
				fmt.Printf("流量查询：您的剩余流量为 %.2f MB\n", flow)
			case 2:
				//2.话费查询
				fmt.Printf("话费查询：您的剩余话费为 %.2f 元\n", balance)
			case 3:
				//3.短信查询
				fmt.Printf("短信查询：您的剩余免费短信为 %d 条\n", shortMessage)
			case 0:
				//返回上一级
				break
			default:
				fmt.Printf("输入错误，没有该选项\n")

			}
		}
	}

}

func phoneChagre() {
	var userChoice int = 1
	for userChoice != 0 {
		chargeMenu()
		fmt.Printf("请输入您的选择：")

		_, inputError := fmt.Scanln(&userChoice)
		//对用户输入校验
		if inputError != nil {
			fmt.Printf("输入错误：%v\n", inputError)
		} else {
			switch userChoice {
			case 1:
				fmt.Printf("请输入充值流量：")
				var userFlow float64
				_, inputError1 := fmt.Scanln(&userFlow)
				if inputError1 != nil {
					fmt.Printf("输入错误：%v\n", inputError1)
				} else {
					//增加用户输入的流量,引用传递才能改变原值
					*(&flow) += userFlow
					fmt.Printf("充值成功！您本次充值流量 %.2f MB，您总的剩余流量为 %.2f MB\n", userFlow, flow)
				}
			case 2:
				fmt.Printf("请输入充值话费：")
				var userBalance float64
				_, inputError1 := fmt.Scanln(&userBalance)
				if inputError1 != nil {
					fmt.Printf("输入错误：%v", inputError1)
				} else {
					//增加用户输入的话费,引用传递才能改变原值
					*(&balance) += userBalance
					fmt.Printf("充值成功！您本次充值话费 %.2f 元，您总的剩余话费为 %.2f 元\n", userBalance, balance)
				}
			case 3:
				fmt.Printf("请输入充值话费：")
				var userMessage int
				_, inputError1 := fmt.Scanln(&userMessage)
				if inputError1 != nil {
					fmt.Printf("输入错误：%v\n", inputError1)
				} else {
					//增加用户输入的短信,引用传递才能改变原值
					*(&shortMessage) += userMessage
					fmt.Printf("充值成功！您本次充值短信 %d 条，您总的剩余短信为 %d 条\n", userMessage, shortMessage)
				}
			case 0:
				break
			default:
				fmt.Printf("输入错误，没有该选项\n")
			}
		}
	}
}

func main() {
	//定义用户输入选项
	var userChoice int = 1

	//循环控制程序不退出，直到用户输入0
	for userChoice != 0 {
		mainMenu()
		fmt.Printf("请输入您的选择：")
		_, inputError := fmt.Scanln(&userChoice)
		//对用户输入校验
		if inputError != nil {
			fmt.Printf("输入错误：%v\n", inputError)
		} else {
			switch userChoice {
			case 1:
				businessQuery()
			case 2:
				phoneChagre()
			case 0:
				fmt.Println("感谢您的使用，再见！")
				break
			}

		}
	}

}
