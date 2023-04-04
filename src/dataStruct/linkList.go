package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

// 在链表尾部插入元素
func (head *ListNode) Insert(value int) {
	newNode := &ListNode{Value: value}
	if head.Next == nil {
		head.Next = newNode
	} else {
		lastNode := head
		for lastNode.Next != nil {
			lastNode = lastNode.Next
		}
		lastNode.Next = newNode
	}
}

// 在链表中查找元素
func (head *ListNode) Search(value int) *ListNode {
	currentNode := head
	for currentNode != nil {
		if currentNode.Value == value {
			return currentNode
		}
		currentNode = currentNode.Next
	}
	return nil
}

// 在链表中删除元素
func (head *ListNode) Delete(value int) {
	if head.Next == nil {
		return
	}
	if head.Next.Value == value {
		head.Next = head.Next.Next
		return
	}
	currentNode := head.Next
	for currentNode.Next != nil {
		if currentNode.Next.Value == value {
			currentNode.Next = currentNode.Next.Next
			return
		}
		currentNode = currentNode.Next
	}
}

// 遍历链表
func (head *ListNode) Traverse() {
	currentNode := head.Next
	for currentNode != nil {
		fmt.Println(currentNode.Value)
		currentNode = currentNode.Next
	}
}

func main() {
	head := &ListNode{}
	head.Insert(1)
	head.Insert(2)
	head.Insert(3)
	head.Insert(4)
	head.Insert(5)

	fmt.Println("Traverse linked list:")
	head.Traverse()

	fmt.Println("Search node:")
	fmt.Println(head.Search(4))
	fmt.Println(head.Search(6))

	fmt.Println("Delete node:")
	head.Delete(3)
	head.Traverse()
}
