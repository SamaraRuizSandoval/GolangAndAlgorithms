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
	{name: "Example 5", input: makeTree(1), want: 1},
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

type TestIsBalanced struct {
	name  string
	input *bst.TreeNode
	want  bool
}

//Example 3: [10,8,12,6,15,4,18]
//         10
//       /   \
//      8     12
//     /       \
//    6         15
//   /           \
//  4             18

var IsBalancedTestCases = []TestIsBalanced{
	{name: "Example 1", input: makeTree(8, 3, 10, 1, 6), want: true},
	{name: "Example 2", input: makeTree(8, 10, 16), want: false},
	{name: "Example 3", input: makeTree(8, 3, 10, 1, 4, 6), want: false},
	{name: "Example 4", input: makeTree(10, 8, 12, 6, 15, 4, 18), want: false},
}

func TestIsBalancedProblem(t *testing.T) {
	for _, tt := range IsBalancedTestCases {
		t.Run(tt.name, func(t *testing.T) {
			got := IsBalanced(tt.input)

			if got != tt.want {
				t.Errorf("got %v but want %v", got, tt.want)
			}
		})
	}
}

// Example 1 [4,2,6,1,3]
//     4
//    / \
//   2   6
//  / \
// 1   3
// Output: 1

// Example 2 [10,5,15,2,7,18]
//     10
//    /  \
//   5   15
//  / \    \
// 2   7   18
// Output: 2

// Example 3 [40,70,50,90]
//   40
//    \
//     70
//    /  \
//   50  90
// Output: 10

var MinimumDiffTestCases = []TestMaximumHeight{
	{name: "Example 1", input: makeTree(4, 2, 6, 1, 3), want: 1},
	{name: "Example 2", input: makeTree(10, 5, 15, 2, 7, 18), want: 2},
	{name: "Example 3", input: makeTree(40, 70, 50, 90), want: 10},
}

func TestMinimumDifference(t *testing.T) {
	for _, tt := range MinimumDiffTestCases {
		t.Run(tt.name, func(t *testing.T) {
			got := MinDiffInBST(tt.input)

			if got != tt.want {
				t.Errorf("got %v but want %v", got, tt.want)
			}
		})
	}
}

type RangeSumInput struct {
	tree  *bst.TreeNode
	Left  int
	Right int
}

type RangeSumTests struct {
	name  string
	input RangeSumInput
	want  int
}

// Example 1 [10,5,15,3,7,18]
// Tree:
//    10
//   /  \
//  5   15
// / \   \
// 3   7   18

// Range: [7, 15]
// Output: 32

// Example 2 [20,5,25,3,10]
// Tree:
//   20
//   /  \
//  5   25
// / \
// 3   10

// Range: [3, 10]
// Output: 18

// Example 3 [30,35,32]
// Tree:
//    30
//      \
//      35
//     /
//    32

// Range: [30, 34]
// Output: 62

var RangeSumCases = []RangeSumTests{
	{name: "Example 1", input: RangeSumInput{tree: makeTree(10, 5, 15, 3, 7, 18), Left: 7, Right: 15}, want: 32},
	{name: "Example 2", input: RangeSumInput{tree: makeTree(20, 5, 25, 3, 10), Left: 3, Right: 10}, want: 18},
	{name: "Example 3", input: RangeSumInput{tree: makeTree(30, 35, 32), Left: 30, Right: 34}, want: 62},
}

func TestRangeSum(t *testing.T) {
	for _, tt := range RangeSumCases {
		t.Run(tt.name, func(t *testing.T) {
			got := RangeSumBST(tt.input.tree, tt.input.Left, tt.input.Right)

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
