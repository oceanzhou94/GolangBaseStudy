/*
《go语言编程从入门到实践》
第二章动手练习：编程实现个人健康评测
BMI计算公式：BMI = 体重（KG） / 身高（M）的平方
BMI < 18.4  过轻
18.5 <=BMI< 24 正常
24 <= BMI< 27 过重
27 <= BMI< 30 轻度肥胖
30 <= BMI< 35 中度肥胖
BMI >= 35 过度肥胖
*/
package main

import (
	"fmt"
)

func main() {
	var name string
	var weight, height float64

	fmt.Printf("please input your name :")
	fmt.Scanln(&name)
	fmt.Printf("please input your weight (KG) :")
	fmt.Scanln(&weight)
	fmt.Printf("please input your height (CM) :")
	fmt.Scanln(&height)

	//print the message and keep two decimals
	fmt.Printf("%s's height: %.2f CM, weight: %.2f KG\n", name, height, weight)

	//unit conversion: CM -> M
	height /= 100

	//caculate the BMI and keep two decimals
	bmi := weight / (height * height)

	//conditional judgment
	if bmi < 18.5 {
		fmt.Printf("your BMI is %.2f, you are too light\n", bmi)
	} else if bmi >= 18.5 && bmi <= 24.0 {
		fmt.Printf("your BMI is %.2f, you are healthy\n", bmi)
	} else if bmi >= 24.0 && bmi <= 27.0 {
		fmt.Printf("your BMI is %.2f, you are a little fat\n", bmi)
	} else if bmi >= 27.0 && bmi <= 30.0 {
		fmt.Printf("your BMI is %.2f, you are light fat\n", bmi)
	} else if bmi >= 30.0 && bmi <= 35.0 {
		fmt.Printf("your BMI is %.2f, you are middle fat\n", bmi)
	} else if bmi >= 35.0 {
		fmt.Printf("your BMI is %.2f, you are very fat\n", bmi)
	}

	fmt.Println("Thanks for using BMI calculate programme")
}
