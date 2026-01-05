package designgurusproblems

import (
	linkedlist "github.com/SamaraRuizSandoval/GolangAndAlgorithms/data_structures/linkedList"
)

//? Given a sorted linked list, remove all the duplicate elements to leave only distinct numbers. The linked list should remain sorted, and the modified list should be returned.
// Input: N = 2
// Output: ["1", "10"]

func DeleteDuplicates(head *linkedlist.ListNode) *linkedlist.ListNode {
	// Initialize the current node as the head of the list.
	current := head

	// Traverse through the list until the end is reached.
	for current != nil && current.Next != nil {
		// If the next node is a duplicate, bypass it.
		if current.Next.Val == current.Val {
			current.Next = current.Next.Next
		} else {
			// If not, move to the next node.
			current = current.Next
		}
	}

	// Return the modified head of the list.
	return head
}
