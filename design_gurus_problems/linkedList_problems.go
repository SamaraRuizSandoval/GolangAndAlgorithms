package designgurusproblems

import (
	linkedlist "github.com/SamaraRuizSandoval/GolangAndAlgorithms/data_structures/linkedList"
)

//? Given a sorted linked list, remove all the duplicate elements to leave only distinct numbers. The linked list should remain sorted, and the modified list should be returned.
// Input: 1 -> 2 -> 2 -> 3
// Output: 1 -> 2 -> 3

func DeleteDuplicates(head *linkedlist.ListNode) *linkedlist.ListNode {
	current := head

	for current != nil && current.Next != nil {
		// If the next node is a duplicate, bypass it.
		if current.Next.Val == current.Val {
			current.Next = current.Next.Next
		} else {
			// If not, move to the next node.
			current = current.Next
		}
	}

	return head
}

//? Given the head of two sorted linked lists, l1 and l2. Return a new sorted list created by merging together the nodes of the first two lists.
// Input: l1 = [1, 3, 5] , l2 [2, 4, 6]
// Output: [1, 2, 3, 4, 5, 6]

func MergeTwoLists(l1 *linkedlist.ListNode, l2 *linkedlist.ListNode) *linkedlist.ListNode {
	newNode := &linkedlist.ListNode{}
	current := newNode

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return newNode.Next
}
