/*
《go语言编程从入门到实践》
第一章动手练习：编程实现人机交互
*/
package main

import "fmt"

func main() {
	var (
		name    string
		age     int
		sex     string
		address string
	)

	fmt.Printf("please input your name: ")
	_, _ = fmt.Scanln(&name)

	fmt.Printf("please input your age: ")
	_, _ = fmt.Scanln(&age)

	fmt.Printf("please input your sex: ")
	_, _ = fmt.Scanln(&sex)

	fmt.Printf("please input your address: ")
	_, _ = fmt.Scanln(&address)

	fmt.Printf("the following is your personal message:\n")
	fmt.Printf("\t   name:  %s\n", name)
	fmt.Printf("\t    age:  %d\n", age)
	fmt.Printf("\t    sex:  %s\n", sex)
	fmt.Printf("\taddress:  %s\n", address)

}
