package bst

import (
	"fmt"
	"io"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BinarySearchTree struct{}

func InorderTraversal(writer io.Writer, root *TreeNode) {
	if root == nil {
		return
	}
	InorderTraversal(writer, root.Left)
	fmt.Fprintf(writer, "%v ", root.Val)
	InorderTraversal(writer, root.Right)

}

func Insert(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return &TreeNode{Val: val}
	}

	if val < node.Val {
		node.Left = Insert(node.Left, val)
	} else if val > node.Val {
		node.Right = Insert(node.Right, val)
	} else {
		//Duplicates not allowed
	}

	return node
}

func Search(node *TreeNode, key int) *TreeNode {
	if node == nil || node.Val == key {
		return node
	}

	fmt.Printf("Checking %v", node.Val)

	if node.Val > key {
		return Search(node.Left, key)
	}

	return Search(node.Right, key)
}
