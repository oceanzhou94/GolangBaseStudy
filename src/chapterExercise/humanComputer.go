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
	fmt.Scanln(&name)

	fmt.Printf("please input your age: ")
	fmt.Scanln(&age)

	fmt.Printf("please input your sex: ")
	fmt.Scanln(&sex)

	fmt.Printf("please input your address: ")
	fmt.Scanln(&address)

	fmt.Printf("the following is your personal message:\n")
	fmt.Printf("\t   name:  %s\n", name)
	fmt.Printf("\t    age:  %d\n", age)
	fmt.Printf("\t    sex:  %s\n", sex)
	fmt.Printf("\taddress:  %s\n", address)

}
