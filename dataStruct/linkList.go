package main

import (
	"fmt"
)

// Node 定义单链表节点结构体
type Node struct {
	data int   // 数据域
	next *Node // 指针域
}

// LinkedList 定义单链表结构体
type LinkedList struct {
	head *Node // 头节点
}

// Add 添加节点到链表末尾
func (l *LinkedList) Add(value int) {
	newNode := &Node{data: value, next: nil}

	// 如果链表为空，则新节点为头节点
	if l.head == nil {
		l.head = newNode
		return
	}

	// 遍历到链表尾部，将新节点加入到末尾
	last := l.head
	for last.next != nil {
		last = last.next
	}
	last.next = newNode
}

// Insert 在指定位置插入节点
func (l *LinkedList) Insert(value int, pos int) {
	newNode := &Node{data: value, next: nil}

	// 如果链表为空，则新节点为头节点
	if l.head == nil {
		l.head = newNode
		return
	}

	// 如果要插入到链表头部，则直接将新节点作为头节点
	if pos == 1 {
		newNode.next = l.head
		l.head = newNode
		return
	}

	// 遍历到指定位置，将新节点插入到链表中
	current := l.head
	var prev *Node
	for i := 1; i < pos; i++ {
		prev = current
		current = current.next
		if current == nil {
			break
		}
	}
	prev.next = newNode
	newNode.next = current
}

// Remove 删除指定节点
func (l *LinkedList) Remove(value int) {
	// 如果链表为空，则直接返回
	if l.head == nil {
		return
	}

	// 如果要删除的节点是头节点
	if l.head.data == value {
		l.head = l.head.next
		return
	}

	// 遍历链表，删除指定节点
	current := l.head
	var prev *Node
	for current != nil && current.data != value {
		prev = current
		current = current.next
	}
	if current != nil {
		prev.next = current.next
	}
}

// Length 获取链表长度
func (l *LinkedList) Length() int {
	length := 0
	current := l.head
	for current != nil {
		length++
		current = current.next
	}
	return length
}

// Print 打印链表
func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	// 创建一个链表对象
	l := LinkedList{}

	// 添加节点
	l.Add(1)
	l.Add(2)
	l.Add(3)

	// 在指定位置插入节点
	l.Insert(4, 1)
	l.Insert(5, 3)

	// 删除指定节点
	l.Remove(3)

	// 打印链表
	l.Print()

	// 获取链表长度
	fmt.Println("链表长度为：", l.Length())
}
