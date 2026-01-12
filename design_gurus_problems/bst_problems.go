package designgurusproblems

import "github.com/SamaraRuizSandoval/GolangAndAlgorithms/data_structures/bst"

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
	// ToDo: Write Your Code Here.
	return false
}
