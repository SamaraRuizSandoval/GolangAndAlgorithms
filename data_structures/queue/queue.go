package queue

import "errors"

type SliceStringQueue struct {
	queue []string
}

type Queue interface {
	Enqueue(item string)
	Dequeue() (string, error)
	Peek() (string, error)
	IsEmpty() bool
}

func NewQueue() *SliceStringQueue {
	return &SliceStringQueue{}
}

var ErrEmptyQueue = errors.New("empty queue")

func (sq *SliceStringQueue) Enqueue(item string) {
	sq.queue = append(sq.queue, item)
}

func (sq *SliceStringQueue) Dequeue() (string, error) {
	if len(sq.queue) == 0 {
		return "", ErrEmptyQueue
	}

	front := sq.queue[0]
	sq.queue = sq.queue[1:]

	return front, nil
}

func (sq *SliceStringQueue) Peek() (string, error) {
	if len(sq.queue) == 0 {
		return "", ErrEmptyQueue
	}

	front := sq.queue[0]

	return front, nil
}

func (sq *SliceStringQueue) IsEmpty() bool {
	return len(sq.queue) == 0
}
