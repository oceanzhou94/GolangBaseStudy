package main

import (
	"errors"
	"fmt"
)

// Stack 定义栈结构体
type Stack struct {
	data []int // 栈的数据存储
	top  int   // 栈顶指针
}

// Init 初始化栈
func (s *Stack) Init() {
	s.data = make([]int, 0)
	s.top = -1
}

// IsEmpty 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

// Top 获取栈顶元素
func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	return s.data[s.top], nil
}

// Push 入栈
func (s *Stack) Push(value int) {
	s.top++
	if s.top >= len(s.data) {
		s.data = append(s.data, value)
	} else {
		s.data[s.top] = value
	}
}

// Pop 出栈
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	value := s.data[s.top]
	s.top--
	return value, nil
}

// Length 获取栈的长度
func (s *Stack) Length() int {
	return s.top + 1
}

// Print 打印栈
func (s *Stack) Print() {
	for i := s.top; i >= 0; i-- {
		fmt.Printf("%d ", s.data[i])
	}
	fmt.Println()
}

func main() {
	// 创建一个栈对象
	s := Stack{}
	s.Init()

	// 入栈
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// 打印栈
	s.Print()

	// 出栈
	value, _ := s.Pop()
	fmt.Printf("Pop %d\n", value)

	// 打印栈
	s.Print()

	// 获取栈顶元素
	value, _ = s.Top()
	fmt.Printf("Top %d\n", value)

	// 获取栈长度
	fmt.Println("Stack length is", s.Length())
}
