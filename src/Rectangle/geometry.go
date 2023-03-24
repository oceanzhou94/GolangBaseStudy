package main

import "awesomeProject/src/Rectangle/rectangle"

func main() {
	//调用其他包文件的函数
	length := 35.8
	width := 18.5
	rectangle.Area(length, width)
}
