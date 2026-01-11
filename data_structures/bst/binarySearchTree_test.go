package bst

import (
	"bytes"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	t.Run("Inorder Traversal", func(t *testing.T) {
		buffer := bytes.Buffer{}

		//         8
		//       /   \
		//      3     10
		//     / \      \
		//    1   6      14
		//       /
		//      4

		tree := makeTree(8, 3, 10, 14, 1, 6, 4)
		InorderTraversal(&buffer, tree)
		got := buffer.String()
		want := "1 3 4 6 8 10 14 "

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("Insert node at root", func(t *testing.T) {
		buffer := bytes.Buffer{}
		var emptyTree *TreeNode
		result := Insert(emptyTree, 5)
		InorderTraversal(&buffer, result)

		got := buffer.String()
		want := "5 "
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("Insert node to left", func(t *testing.T) {
		buffer := bytes.Buffer{}
		root := &TreeNode{Val: 10}
		result := Insert(root, 5)
		InorderTraversal(&buffer, result)

		got := buffer.String()
		want := "5 10 "
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("Insert node to right", func(t *testing.T) {
		buffer := bytes.Buffer{}
		root := &TreeNode{Val: 10}
		result := Insert(root, 18)
		InorderTraversal(&buffer, result)

		got := buffer.String()
		want := "10 18 "
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("Search element that exists", func(t *testing.T) {
		tree := makeTree(8, 3, 10, 14, 1, 6, 4)
		result := Search(tree, 6)

		if result == nil {
			t.Errorf("got nil but node exists")

		}
	})

	t.Run("Search element that doesn't exists", func(t *testing.T) {
		tree := makeTree(8, 3, 10, 14, 1, 6, 4)
		result := Search(tree, 87)

		if result != nil {
			t.Errorf("got nbde but node doesn't exists")

		}
	})
}

func makeTree(vals ...int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}

	root := &TreeNode{Val: vals[0]}
	current := root

	for _, v := range vals[1:] {
		Insert(current, v)
	}
	return root
}
