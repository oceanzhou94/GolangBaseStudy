/*
结构体变量
理解结构体的声明与使用
*/
package main

import (
	"fmt"
)

//定义结构体person
type person struct {
	name string
	age  int
	sex  string
}

//结构体嵌套
type student struct {
	basicMessage person
	studentId    string
	score        float64
}

//定义结构体的方法，属于结构体person
func (p *person) getInfo() {
	fmt.Println("the information of perosn:")
	fmt.Printf("\tname is %s\n", p.name)
	fmt.Printf("\tage is %d\n", p.age)
	fmt.Printf("\tsex is %s\n", p.sex)
}

//定义结构体方法，属于结构体person
func (p *person) code() {
	fmt.Printf("%s is coding\n", p.name)
}

//定义结构体方法，属于结构体person
func (p *person) read(book string) {
	fmt.Printf("%s is reaing book 《%s》\n", p.name, book)
}

//自定义定义结构体构造函数
func personConstructor(name string, age int, sex string) *person {
	p := person{
		name: name,
		age:  age,
		sex:  sex,
	}
	return &p
}

//结构体方法，属于student结构体
func (s *student) getInfo() {
	fmt.Println("the information of student:")
	fmt.Printf("basicMessage: name is %s, age is %d, sex is %s\n",
		s.basicMessage.name, s.basicMessage.age, s.basicMessage.sex)
	fmt.Printf("student information: student id is %s, score is %.2f\n",
		s.studentId, s.score)
}

//定义结构体方法，属于结构体person
func (s *student) read(book string) {
	fmt.Printf("%s is reaing book 《%s》\n", s.basicMessage.name, book)
}

func main() {
	//初始化结构体变量p1
	var p1 person = person{
		name: "person1",
		age:  18,
		sex:  "男",
	}
	//调用结构体方法
	p1.getInfo()

	//修改结构体成员变量
	p1.name = "person0"
	p1.sex = "女"
	p1.getInfo()

	//使用构造函数创建结构体
	p2 := personConstructor("person2", 19, "男")
	p2.getInfo()

	//调用结构体方法code()、read()
	book := "Go语言编程从入门到实践"
	p3 := personConstructor("person3", 20, "女")
	p3.getInfo()
	p3.code()
	p3.read(book)

	//初始化一个student结构体
	s0 := student{
		basicMessage: person{name: "xiaoMei", age: 17, sex: "女"},
		studentId:    "0933",
		score:        599.95,
	}

	//调用结构体方法
	s0.getInfo()
	s0.read("Go语言进阶")

}
