package linkedlist

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var ErrInvalidPosition = errors.New("invalid linkedList position")
var ErrListIsEmpty = errors.New("list is empty")

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
	Next *DoubleListNode
	Prev *DoubleListNode
}

type DoubleLinkedList struct {
	Head *DoubleListNode
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
		fmt.Fprintf(writer, "%v ⇄ ", current.Val)
		current = current.Next
	}

}

func (list *DoubleLinkedList) TraverseBackward(writer io.Writer) {
	current := list.Head

	if list.Head == nil {
		return
	}

	for current.Next != nil {
		current = current.Next
	}

	for current != nil {
		fmt.Fprintf(writer, "%v ⇄ ", current.Val)
		current = current.Prev
	}

}

func (list *DoubleLinkedList) InsertAtPosition(val, position int) error {
	newNode := &DoubleListNode{Val: val, Next: nil, Prev: nil}

	if position == 0 {
		newNode.Next = list.Head

		if list.Head != nil {
			list.Head.Prev = newNode
		}

		list.Head = newNode
		return nil
	}

	current := list.Head
	for i := 1; i <= position-1 && current != nil; i++ {
		current = current.Next
	}

	if current == nil {
		return ErrInvalidPosition
	}

	newNode.Next = current.Next
	if current.Next != nil {
		current.Next.Prev = newNode
	}
	current.Next = newNode
	newNode.Prev = current
	return nil
}

func (list *DoubleLinkedList) DeleteAtPosition(position int) error {
	if list.Head == nil {
		return ErrInvalidPosition
	}

	if position == 0 {
		if list.Head.Next == nil {
			list.Head = nil
			return nil
		}

		list.Head = list.Head.Next
		list.Head.Prev = nil
		return nil
	}

	current := list.Head
	for i := 0; i < position-1 && current != nil; i++ {
		current = list.Head.Next
	}

	if current == nil || current.Next == nil {
		return ErrInvalidPosition
	}

	toDeleteNode := current.Next
	current.Next = toDeleteNode.Next
	if toDeleteNode.Next != nil {
		toDeleteNode.Next.Prev = current
	}

	return nil

}
