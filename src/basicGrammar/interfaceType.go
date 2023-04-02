/*
接口类型，用于封装具有共性的一系列方法
通过不同的结构体实现接口方法从而实现多态
*/

package main

import "fmt"

//定义接口--动作
type actions interface {
	walk() //行走
	run()  //跑动
	eat()  //进食
}

//定义结构体--狗
type dog struct {
	dogType  string //种类
	dogColor string //毛发颜色
}

//定义结构体--猫
type cat struct {
	catType  string //种类
	catColor string //毛发颜色
}

//结构体dog实现接口方法--walk
func (d *dog) walk() {
	fmt.Println("walk：dog is walking!")
	fmt.Printf("it's a dog,its type is %s,its color is %s\n", d.dogType, d.dogColor)
}

//结构体dog实现接口方法--run
func (d *dog) run() {
	fmt.Printf("run：a %s dog is running, speed is fast!\n", d.dogType)
}

//结构体dog实现接口方法--eat
func (d *dog) eat() {
	fmt.Printf("eat：a dog is eating, its color is %s\n", d.dogColor)
}

//结构体cat实现接口方法--walk
func (c *cat) walk() {
	fmt.Println("walk：cat is walking!")
	fmt.Printf("it's a cat,its type is %s,its color is %s\n", c.catType, c.catColor)
}

//结构体cat实现接口方法--run
func (c *cat) run() {
	fmt.Printf("run：a %s cat is running, speed is fast!\n", c.catType)
}

//结构体cat实现接口方法--eat
func (c *cat) eat() {
	fmt.Printf("eat：cat is eating, its color is %s\n", c.catColor)
}

func main() {
	//接口实例
	var dogAct, catAct actions

	//结构体实例赋值给接口实例
	dogAct = &dog{"金毛", "金黄色"}
	//调用接口方法
	dogAct.walk()
	dogAct.run()
	dogAct.eat()

	//结构体实例赋值给接口实例
	catAct = &cat{"蓝猫", "蓝灰色"}
	//调用接口方法
	catAct.walk()
	catAct.run()
	catAct.eat()

}
