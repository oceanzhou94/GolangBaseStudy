package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// 插入节点
func (node *TreeNode) Insert(value int) {
	if node == nil {
		return
	}
	if value < node.Value {
		if node.Left == nil {
			node.Left = &TreeNode{Value: value}
		} else {
			node.Left.Insert(value)
		}
	} else {
		if node.Right == nil {
			node.Right = &TreeNode{Value: value}
		} else {
			node.Right.Insert(value)
		}
	}
}

// 查找节点
func (node *TreeNode) Search(value int) bool {
	if node == nil {
		return false
	}
	if value < node.Value {
		return node.Left.Search(value)
	} else if value > node.Value {
		return node.Right.Search(value)
	} else {
		return true
	}
}

// 删除节点
func (node *TreeNode) Delete(value int) *TreeNode {
	if node == nil {
		return nil
	}
	if value < node.Value {
		node.Left = node.Left.Delete(value)
	} else if value > node.Value {
		node.Right = node.Right.Delete(value)
	} else {
		if node.Left == nil && node.Right == nil {
			return nil
		} else if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		} else {
			minNode := node.Right.MinNode()
			node.Value = minNode.Value
			node.Right = node.Right.Delete(minNode.Value)
		}
	}
	return node
}

// 遍历二叉树
func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	fmt.Println(node.Value)
	node.Right.Traverse()
}

// 获取最小节点
func (node *TreeNode) MinNode() *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func main() {
	root := &TreeNode{Value: 5}
	root.Insert(3)
	root.Insert(7)
	root.Insert(1)
	root.Insert(9)
	root.Insert(2)
	root.Insert(4)

	fmt.Println("Traverse tree:")
	root.Traverse()

	fmt.Println("Search node:")
	fmt.Println(root.Search(4))
	fmt.Println(root.Search(8))

	fmt.Println("Delete node:")
	root = root.Delete(3)
	root.Traverse()
}
