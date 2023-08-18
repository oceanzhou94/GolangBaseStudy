package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("终于，配置完成了这个破玩意儿！")

	var input int
	var input2 int
	n, e := fmt.Scanln(&input, &input2)
	fmt.Println("n", n)
	fmt.Println("e", e)
}
