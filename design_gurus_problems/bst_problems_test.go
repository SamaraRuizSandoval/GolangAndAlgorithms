package designgurusproblems

import (
	"testing"

	"github.com/SamaraRuizSandoval/GolangAndAlgorithms/data_structures/bst"
)

type TestMaximumHeight struct {
	name  string
	input *bst.TreeNode
	want  int
}

//Example 1: [8,3,10,1,6]
//         8
//       /   \
//      3     10
//     / \
//   1    6

//Example 2: [8,10,6]
//         8
//           \
//            10
//      		\
//        	     16

//Example 3: [8,3,10,1,4,6]
//         8
//       /   \
//      3     10
//     / \
// 	  1	  4
//         \
//			6

var maximumHeightTestCases = []TestMaximumHeight{
	{name: "Example 1", input: makeTree(8, 3, 10, 1, 6), want: 3},
	{name: "Example 2", input: makeTree(8, 10, 16), want: 3},
	{name: "Example 3", input: makeTree(8, 3, 10, 1, 4, 6), want: 4},
	{name: "Example 4", input: makeTree(), want: 0},
	{name: "Example 4", input: makeTree(1), want: 1},
}

func TestMaximumHeightProblem(t *testing.T) {
	for _, tt := range maximumHeightTestCases {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxDepth(tt.input)

			if got != tt.want {
				t.Errorf("got %v but want %v", got, tt.want)
			}
		})
	}
}

func makeTree(vals ...int) *bst.TreeNode {
	if len(vals) == 0 {
		return nil
	}

	root := &bst.TreeNode{Val: vals[0]}
	current := root

	for _, v := range vals[1:] {
		bst.Insert(current, v)
	}
	return root
}
