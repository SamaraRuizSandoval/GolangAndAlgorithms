package dgQueueProblems

import (
	"github.com/SamaraRuizSandoval/GolangAndAlgorithms/data_structures/queue"
)

//? Given an integer N, generate all binary numbers from 1 to N and return them as a list of strings.
// Input: N = 2
// Output: ["1", "10"]

func GenerateBinaryNumbers(n int) []string {
	res := make([]string, n)

	queue := queue.NewQueue()
	queue.Enqueue("1")

	for i := 0; i < n; i++ {
		binaryNum, _ := queue.Dequeue()
		res[i] = binaryNum

		queue.Enqueue(binaryNum + "0")
		queue.Enqueue(binaryNum + "1")

	}

	return res
}

//? Given an integer array arr and an integer k, return the result list containing the maximum for each and every contiguous subarray of size k.
// In other words, result[i] = max(arr[0],..., arr[k]), result[1] = max(arr[1],...arr[k+1]), etc.

// Input: arr = [1, 2, 3, 1, 4, 5, 2, 3, 6], k = 3
// Output: [3, 3, 4, 5, 5, 5, 6]
// Description: Here, subarray [1,2,3] has maximum 3, [2,3,1] has maximum 3, [3,1,4] has maximum 4, [1,4,5] has maximum 5, [4,5,2] has maximum 5, [5,2,3] has maximum 5, and [2,3,6] has maximum 6.

func PrintMax(arr []int, k int) []int {
	result := []int{}
	// ToDo
	return result
}

//? Given a string s, determine if that string is a palindrome using a queue data structure. Return true if that string is a palindrome, otherwise return false
// Input: s = "madam"
// Output: true

func CheckPalindrome(str string) bool {
	// ToDo: We need a double sided queue
	return false
}
