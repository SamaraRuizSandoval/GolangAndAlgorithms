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

	if node.Val > key {
		return Search(node.Left, key)
	}

	return Search(node.Right, key)
}

func DeleteNode(node *TreeNode, key int) *TreeNode {
	if node == nil {
		return nil
	}

	if key < node.Val {
		node.Left = DeleteNode(node.Left, key)
	} else if key > node.Val {
		node.Right = DeleteNode(node.Right, key)
	} else { // Node found, handle deletion
		// Case 1: Node has no left child
		if node.Left == nil {
			return node.Right
		}
		// Case 2: Node has no right child
		if node.Right == nil {
			return node.Left
		}

		// Case 3: Node has two children
		successor := minValueNode(node.Right)              // Find inorder successor
		node.Val = successor.Val                           // Copy the successor's value
		node.Right = DeleteNode(node.Right, successor.Val) // Delete the successor
	}

	return node
}

// Helper method to find the inorder successor (smallest node in right subtree)
func minValueNode(node *TreeNode) *TreeNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}
