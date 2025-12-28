package main

import (
	"os"
	"time"

	linkedlist "github.com/SamaraRuizSandoval/GolangAndAlgorithms/data_structures/linkedList"
)

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	list := linkedlist.SinglyLinkedList{}

	// Creating linked list: 10 → 20 → 30 → 40 → NULL
	list.Head = &linkedlist.ListNode{Val: 10}
	list.Head.Next = &linkedlist.ListNode{Val: 20}
	list.Head.Next.Next = &linkedlist.ListNode{Val: 30}
	list.Head.Next.Next.Next = &linkedlist.ListNode{Val: 40}

	list.Traverse(os.Stdout)
}
