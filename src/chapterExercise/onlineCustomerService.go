/*编程实现模拟在线客服*/
package main

import (
	"fmt"
	"time"
)

var balance float64 = 7.55 //话费余额
var flow float64 = 1024.88 // 流量
var shortMessage int = 49  //短信
//手机业务键值对
var business map[string]string = map[string]string{
	"飞信业务": "未办理",
	"黄钻业务": "未办理",
	"宽带业务": "已办理",
	"5G通讯": "已办理",
	"国内漫游": "未办理",
}

//功能菜单
func mainMenu() {
	fmt.Println("--------在线客服模拟--------")
	fmt.Println("\t业务查询请按1")
	fmt.Println("\t手机充值请按2")
	fmt.Println("\t业务办理请按3")
	fmt.Println("\t语音导航请按4")
	fmt.Println("\t人工服务请按5")
	fmt.Println("\t退出服务请按0")
}

//业务查询二级菜单
func queryMenu() {
	fmt.Println("--------业务查询二级菜单--------")
	fmt.Println("\t流量查询请按1")
	fmt.Println("\t话费查询请按2")
	fmt.Println("\t短信查询请按3")
	fmt.Println("\t已办业务请按4")
	fmt.Println("\t返回上级请按0")
}

//手机充值二级菜单
func chargeMenu() {
	fmt.Println("--------手机充值二级菜单--------")
	fmt.Println("\t流量充值请按1")
	fmt.Println("\t话费充值请按2")
	fmt.Println("\t短信充值请按3")
	fmt.Println("\t返回上级请按0")
}

//业务办理二级菜单
func businessHandingMenu() {
	fmt.Println("--------业务办理二级菜单--------")
	fmt.Println("\t飞信业务请按1")
	fmt.Println("\t黄钻业务请按2")
	fmt.Println("\t宽带办理请按3")
	fmt.Println("\t5G 业务请按4")
	fmt.Println("\t国内漫游请按5")
	fmt.Println("\t返回上级请按0")
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
				//睡眠一秒
				time.Sleep(1 * time.Second)
			case 2:
				//2.话费查询
				fmt.Printf("话费查询：您的剩余话费为 %.2f 元\n", balance)
				//睡眠一秒
				time.Sleep(1 * time.Second)
			case 3:
				//3.短信查询
				fmt.Printf("短信查询：您的剩余免费短信为 %d 条\n", shortMessage)
				//睡眠一秒
				time.Sleep(1 * time.Second)
			case 4:
				//4.已办业务查询
				fmt.Printf("业务查询：您已经办理的业务有 ")
				for key, value := range business {
					if value == "已办理" {
						fmt.Printf("%s  ", key)
					}
					time.Sleep(1 * time.Second) //睡眠一秒
				}
				fmt.Println()

				//未办理业务
				fmt.Printf("你还能办理以下业务：")
				for key, value := range business {
					if value == "未办理" {
						fmt.Printf("%s  ", key)
					}
					time.Sleep(1 * time.Second) //睡眠一秒
				}
				fmt.Println()
			case 0:
				//返回上一级
				//睡眠一秒
				time.Sleep(1 * time.Second)
				break
			default:
				fmt.Printf("输入错误，没有该选项\n")
			}
		}
	}

}

//手机充值逻辑
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
					//睡眠一秒
					time.Sleep(1 * time.Second)
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
					//睡眠一秒
					time.Sleep(1 * time.Second)
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
					//睡眠一秒
					time.Sleep(1 * time.Second)
				}
			case 0:
				//返回上一级菜单
				//睡眠一秒
				time.Sleep(1 * time.Second)
				break
			default:
				fmt.Printf("输入错误，没有该选项\n")
			}
		}
	}
}

//业务办理逻辑
func businessHandling() {
	//提示已办理业务
	fmt.Printf("您的账户已经办理以下业务：")
	for key, value := range business {
		if value == "已办理" {
			fmt.Printf("%s  ", key)
		}
		time.Sleep(1 * time.Second) //睡眠一秒
	}
	fmt.Println()
	var userChoice int = 1
	for userChoice != 0 {
		businessHandingMenu()
		fmt.Printf("请输入您的选择：")

		_, inputError := fmt.Scanln(&userChoice)
		//对用户输入校验
		if inputError != nil {
			fmt.Printf("输入错误：%v\n", inputError)
		} else {
			switch userChoice {
			case 1:
				if business["飞信业务"] == "已办理" {
					//睡眠一秒
					time.Sleep(1 * time.Second)
					fmt.Printf("业务办理失败，您已办理该业务，无需重复办理\n")
				} else {
					//办理成功，修改键值对
					fmt.Println("恭喜你，业务办理成功！")
					business["飞信业务"] = "已办理"
					//睡眠一秒
					time.Sleep(1 * time.Second)
				}
			case 2:
				if business["黄钻业务"] == "已办理" {
					//睡眠一秒
					time.Sleep(1 * time.Second)
					fmt.Printf("业务办理失败，您已办理该业务，无需重复办理\n")
				} else {
					//办理成功，修改键值对
					fmt.Println("恭喜你，业务办理成功！")
					business["黄钻业务"] = "已办理"
					//睡眠一秒
					time.Sleep(1 * time.Second)
				}
			case 3:
				if business["宽带业务"] == "已办理" {
					//睡眠一秒
					time.Sleep(1 * time.Second)
					fmt.Printf("业务办理失败，您已办理该业务，无需重复办理\n")
				} else {
					//办理成功，修改键值对
					fmt.Println("恭喜你，业务办理成功！")
					business["宽带业务"] = "已办理"
					//睡眠一秒
					time.Sleep(1 * time.Second)
				}
			case 4:
				if business["5G通讯"] == "已办理" {
					//睡眠一秒
					time.Sleep(1 * time.Second)
					fmt.Printf("业务办理失败，您已办理该业务，无需重复办理\n")
				} else {
					//办理成功，修改键值对
					fmt.Println("恭喜你，业务办理成功！")
					business["5G通讯"] = "已办理"
					//睡眠一秒
					time.Sleep(1 * time.Second)
				}
			case 5:
				if business["国内漫游"] == "已办理" {
					//睡眠一秒
					time.Sleep(1 * time.Second)
					fmt.Printf("业务办理失败，您已办理该业务，无需重复办理\n")
				} else {
					//办理成功，修改键值对
					fmt.Println("恭喜你，业务办理成功！")
					business["国内漫游"] = "已办理"
					//睡眠一秒
					time.Sleep(1 * time.Second)
				}
			case 0:
				//返回上一级
				//睡眠一秒
				time.Sleep(1 * time.Second)
				break

			default:
				fmt.Printf("输入错误，没有该选项\n")
			}
		}
	}
}

//人工服务逻辑
func manualService() {
	fmt.Println("请稍等，正在转接人工服务")
	for i := 0; i < 6; i++ {
		fmt.Printf(".")
		time.Sleep(1 * time.Second)
	}
	fmt.Println()
	fmt.Println("人工服务忙，请稍后重试！")
}

//语音导航逻辑
func audioGuide() {
	var destination string
loop:
	fmt.Printf("请输入您的目的地：")
	_, err := fmt.Scanln(&destination)
	if err != nil {
		fmt.Println("目的地输入有误，请重新输入")
		goto loop
	} else {
		fmt.Printf("正在为您导航至：%s，祝您旅途愉快！\n", destination)
	}
	//睡眠一秒
	time.Sleep(1 * time.Second)
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
				//睡眠一秒
				time.Sleep(1 * time.Second)
				businessQuery()
			case 2:
				//睡眠一秒
				time.Sleep(1 * time.Second)
				phoneChagre()
			case 3:
				//睡眠一秒
				time.Sleep(1 * time.Second)
				businessHandling()
			case 4:
				//睡眠一秒
				time.Sleep(1 * time.Second)
				audioGuide()
			case 5:
				//睡眠一秒
				time.Sleep(1 * time.Second)
				manualService()
			case 0:
				//睡眠一秒
				time.Sleep(1 * time.Second)
				fmt.Println("感谢您的使用，再见！")
				break
			default:
				//睡眠一秒
				time.Sleep(1 * time.Second)
				fmt.Printf("输入错误，没有该选项\n")
			}

		}
	}

}
