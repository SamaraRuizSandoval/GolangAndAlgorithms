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

type InputLists struct {
	l1 linkedlist.ListNode
	l2 linkedlist.ListNode
}

type MergeTwoListsTest struct {
	name  string
	input InputLists
	want  string
}

var testMergeTwoLists = []MergeTwoListsTest{
	{name: "Example 1", input: InputLists{*makeList(1, 3, 5), *makeList(2, 4, 6)}, want: "1 → 2 → 3 → 4 → 5 → 6 → "},
	{name: "Example 2", input: InputLists{*makeList(1), *makeList(2, 3, 4)}, want: "1 → 2 → 3 → 4 → "},
	{name: "Example 3", input: InputLists{*makeList(2, 4, 6), *makeList(1, 3, 5)}, want: "1 → 2 → 3 → 4 → 5 → 6 → "},
}

func TestMergeTwoLists(t *testing.T) {
	for _, tt := range testMergeTwoLists {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeTwoLists(&tt.input.l1, &tt.input.l2)
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

func makeList(vals ...int) *linkedlist.ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &linkedlist.ListNode{Val: vals[0]}
	curr := head
	for _, v := range vals[1:] {
		curr.Next = &linkedlist.ListNode{Val: v}
		curr = curr.Next
	}
	return head
}
