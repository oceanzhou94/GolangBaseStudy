package main

import (
	"fmt"
)

// BinaryTree 二叉树结构体
type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

// Insert 插入节点
func (bt *BinaryTree) Insert(val int) *BinaryTree {
	// 如果树为空，则创建一个节点并返回
	if bt == nil {
		return &BinaryTree{Val: val}
	}

	// 如果插入的值小于当前节点，则将其插入左子树
	if val < bt.Val {
		bt.Left = bt.Left.Insert(val)
		return bt
	}

	// 如果插入的值大于等于当前节点，则将其插入右子树
	bt.Right = bt.Right.Insert(val)
	return bt
}

// Search 查找节点
func (bt *BinaryTree) Search(val int) *BinaryTree {
	// 如果节点为空，则返回 nil
	if bt == nil {
		return nil
	}

	// 如果查找的值等于当前节点的值，则返回当前节点
	if bt.Val == val {
		return bt
	}

	// 如果查找的值小于当前节点的值，则在左子树中查找
	if val < bt.Val {
		return bt.Left.Search(val)
	}

	// 如果查找的值大于当前节点的值，则在右子树中查找
	return bt.Right.Search(val)
}

// PreorderTraversal 前序遍历
func (bt *BinaryTree) PreorderTraversal() []int {
	if bt == nil {
		return []int{}
	}

	result := []int{bt.Val}
	result = append(result, bt.Left.PreorderTraversal()...)
	result = append(result, bt.Right.PreorderTraversal()...)

	return result
}

// InorderTraversal 中序遍历
func (bt *BinaryTree) InorderTraversal() []int {
	if bt == nil {
		return []int{}
	}

	result := bt.Left.InorderTraversal()
	result = append(result, bt.Val)
	result = append(result, bt.Right.InorderTraversal()...)

	return result
}

// PostorderTraversal 后序遍历
func (bt *BinaryTree) PostorderTraversal() []int {
	if bt == nil {
		return []int{}
	}

	result := bt.Left.PostorderTraversal()
	result = append(result, bt.Right.PostorderTraversal()...)
	result = append(result, bt.Val)

	return result
}

func main() {
	bt := &BinaryTree{Val: 5}
	bt.Insert(3)
	bt.Insert(8)
	bt.Insert(1)
	bt.Insert(4)
	bt.Insert(7)
	bt.Insert(9)

	fmt.Println("前序遍历：", bt.PreorderTraversal())
	fmt.Println("中序遍历：", bt.InorderTraversal())
	fmt.Println("后序遍历：", bt.PostorderTraversal())

	node := bt.Search(7)
	if node == nil {
		fmt.Println("未找到节点")
	} else {
		fmt.Println("找到节点：", node.Val)
	}
}
