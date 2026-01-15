package designgurusproblems

import (
	"math"

	"github.com/SamaraRuizSandoval/GolangAndAlgorithms/data_structures/bst"
)

//? Given a root node of the binary tree, return the depth (or height) of a binary tree.
//The Depth of the binary tree refers to the number of nodes along the longest path from the root node to the farthest leaf
// node. If the tree is empty, the depth is 0.

//Input: [8,3,10,1,6]
//         8
//       /   \
//      3     10
//     / \
//    1   6

//Expected Output: 3

func MaxDepth(root *bst.TreeNode) int {
	if root == nil {
		return 0
	}

	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)

	max := left
	if right > left {
		max = right
	}

	return 1 + max
}

//? Determine if a binary tree is height-balanced.

func IsBalanced(root *bst.TreeNode) bool {
	return depth(root) != -1
}

// depth is a helper function to get the depth of the tree from a given node
func depth(node *bst.TreeNode) int {
	if node == nil {
		return 0
	}

	leftDepth := depth(node.Left)
	// If the left subtree is unbalanced, we return -1 to indicate it's not balanced
	if leftDepth == -1 {
		return -1
	}

	rightDepth := depth(node.Right)
	// If the right subtree is unbalanced, we return -1 to indicate it's not balanced
	if rightDepth == -1 {
		return -1
	}

	// If the current node is unbalanced, return -1 to indicate it's not balanced
	if math.Abs(float64(leftDepth-rightDepth)) > 1 {
		return -1
	}

	// If it's balanced, we return the depth of the current subtree
	return int(math.Max(float64(leftDepth), float64(rightDepth))) + 1
}

//? Given a Binary Search Tree (BST), you are required to find the smallest difference between the values of any two different nodes.

var minDiff int
var previous *int

func MinDiffInBST(root *bst.TreeNode) int {
	minDiff = root.Val
	previous = nil
	ReverseInOrderTraversal(root)
	return minDiff
}

func ReverseInOrderTraversal(node *bst.TreeNode) {
	if node == nil {
		return
	}
	ReverseInOrderTraversal(node.Right)
	if previous != nil {
		diff := *previous - node.Val
		if diff < minDiff {
			minDiff = diff
		}
	}
	previous = &node.Val
	ReverseInOrderTraversal(node.Left)
}
