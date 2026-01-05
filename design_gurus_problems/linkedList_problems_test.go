package designgurusproblems

import (
	"bytes"
	"testing"

	linkedlist "github.com/SamaraRuizSandoval/GolangAndAlgorithms/data_structures/linkedList"
)

type RemoveDuplicates struct {
	name  string
	input linkedlist.ListNode
	want  string
}

var exampleOne = linkedlist.ListNode{Val: 1, Next: &linkedlist.ListNode{Val: 1, Next: &linkedlist.ListNode{Val: 2, Next: nil}}}
var exampleTwo = linkedlist.ListNode{Val: 1, Next: &linkedlist.ListNode{Val: 2, Next: &linkedlist.ListNode{Val: 2, Next: &linkedlist.ListNode{Val: 3, Next: nil}}}}
var exampleThree = linkedlist.ListNode{Val: 3, Next: &linkedlist.ListNode{Val: 3, Next: &linkedlist.ListNode{Val: 3, Next: nil}}}
var exampleFour = linkedlist.ListNode{Val: 1, Next: &linkedlist.ListNode{Val: 1, Next: &linkedlist.ListNode{Val: 2, Next: &linkedlist.ListNode{Val: 3, Next: &linkedlist.ListNode{Val: 4, Next: &linkedlist.ListNode{Val: 4, Next: nil}}}}}}

var testRemoveDuplicates = []RemoveDuplicates{
	{name: "Example 1", input: exampleOne, want: "1 → 2 → "},
	{name: "Example 2", input: exampleTwo, want: "1 → 2 → 3 → "},
	{name: "Example 3", input: exampleThree, want: "3 → "},
	{name: "Example 4", input: exampleFour, want: "1 → 2 → 3 → 4 → "},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, tt := range testRemoveDuplicates {
		t.Run(tt.name, func(t *testing.T) {
			result := DeleteDuplicates(&tt.input)
			list := linkedlist.SinglyLinkedList{Head: result}

			buffer := bytes.Buffer{}
			list.Traverse(&buffer)

			got := buffer.String()

			if got != tt.want {
				t.Errorf("got %q but want %q", got, tt.want)
			}

		})
	}
}
