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

var testRemoveDuplicates = []RemoveDuplicates{
	{name: "Example 1", input: *makeList(1, 1, 2), want: "1 → 2 → "},
	{name: "Example 2", input: *makeList(1, 2, 2, 3), want: "1 → 2 → 3 → "},
	{name: "Example 3", input: *makeList(3, 3, 3), want: "3 → "},
	{name: "Example 4", input: *makeList(1, 1, 2, 3, 4, 4), want: "1 → 2 → 3 → 4 → "},
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

type IsPalindromeTest struct {
	name  string
	input linkedlist.DoubleListNode
	want  bool
}

var testIsPalindrome = []IsPalindromeTest{
	{name: "Example 1", input: *makeDoubledLinkedList(1, 2, 3, 2, 1), want: true},
	{name: "Example 2", input: *makeDoubledLinkedList(1, 2, 2, 3), want: false},
	{name: "Example 3", input: *makeDoubledLinkedList(1, 1, 1, 1), want: true},
}

func TestIsPalindrome(t *testing.T) {
	for _, tt := range testIsPalindrome {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(&tt.input)

			if result != tt.want {
				t.Errorf("got %v but want %v", result, tt.want)
			}

		})
	}
}

func makeDoubledLinkedList(vals ...int) *linkedlist.DoubleListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &linkedlist.DoubleListNode{Val: vals[0], Prev: nil}
	curr := head
	for _, v := range vals[1:] {
		curr.Next = &linkedlist.DoubleListNode{Val: v}
		curr.Next.Prev = curr
		curr = curr.Next
	}
	return head
}

var testSwapPairs = []RemoveDuplicates{
	{name: "Example 1", input: *makeList(1, 2, 3, 4), want: "2 → 1 → 4 → 3 → "},
	{name: "Example 2", input: *makeList(7, 8, 9, 10, 11), want: "8 → 7 → 10 → 9 → 11 → "},
	{name: "Example 3", input: *makeList(5, 6), want: "6 → 5 → "},
}

func TestSwapPairs(t *testing.T) {
	for _, tt := range testSwapPairs {
		t.Run(tt.name, func(t *testing.T) {
			result := SwapPairs(&tt.input)
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
