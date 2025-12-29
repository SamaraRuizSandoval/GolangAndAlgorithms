package linkedlist

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var ErrInvalidPosition = errors.New("invalid linkedList position")

type ListNode struct {
	Val  int
	Next *ListNode
}

type SinglyLinkedList struct {
	Head *ListNode
}

// Double Linked List

type DoubleListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode
}

type DoubleLinkedList struct {
	Head *ListNode
}

func (list *SinglyLinkedList) Traverse(writer io.Writer) {
	current := list.Head

	for current != nil {
		fmt.Fprintf(writer, "%v → ", current.Val)
		current = current.Next
	}

}

func (list *SinglyLinkedList) InsertAtPosition(val, position int) error {
	newNode := ListNode{Val: val, Next: nil}
	current := list.Head

	if position == 0 {
		newNode.Next = list.Head
		list.Head = &newNode
		return nil
	}

	for i := 1; i <= position-1 && current != nil; i++ {
		current = current.Next
	}

	if current == nil {
		return ErrInvalidPosition
	}

	newNode.Next = current.Next
	current.Next = &newNode

	return nil
}

func (list *SinglyLinkedList) DeleteAtPosition(position int) error {
	current := list.Head
	list.Traverse(os.Stdout)
	if list.Head == nil {
		return ErrInvalidPosition
	}

	if position == 0 {
		list.Head = list.Head.Next
		return nil
	}

	for i := 0; i < position-1 && current != nil; i++ {
		current = list.Head.Next
	}

	if current == nil || current.Next == nil {
		return ErrInvalidPosition
	}

	current.Next = current.Next.Next

	return nil
}

func (list *DoubleLinkedList) TraverseForward(writer io.Writer) {
	current := list.Head

	for current != nil {
		fmt.Fprintf(writer, "%v → ", current.Val)
		current = current.Next
	}

}

func (list *DoubleLinkedList) TraverseBackward(writer io.Writer) {
	current := list.Head

	for current != nil {
		fmt.Fprintf(writer, "%v → ", current.Val)
		current = current.Next
	}

}
