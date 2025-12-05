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
