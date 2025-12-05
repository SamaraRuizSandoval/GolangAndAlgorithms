package queue

import "errors"

type SliceQueue struct {
	queue []int
}

type Queue interface {
	Enqueue(item int)
	Dequeue() (int, error)
	Peek() (int, error)
	IsEmpty() bool
}

var ErrEmptyQueue = errors.New("empty queue")

func (sq *SliceQueue) Enqueue(item int) {
	sq.queue = append(sq.queue, item)
}

func (sq *SliceQueue) Dequeue() (int, error) {
	if len(sq.queue) == 0 {
		return 0, ErrEmptyQueue
	}

	front := sq.queue[0]
	sq.queue = sq.queue[1:]

	return front, nil
}

func (sq *SliceQueue) Peek() (int, error) {
	if len(sq.queue) == 0 {
		return 0, ErrEmptyQueue
	}

	front := sq.queue[0]

	return front, nil
}

func (sq *SliceQueue) IsEmpty() bool {
	return len(sq.queue) == 0
}
