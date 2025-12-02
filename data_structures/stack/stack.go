package stack

import (
	"errors"
	"fmt"
)

type SliceStack struct {
	stack []int
}

var ErrEmptyStack = errors.New("empty stack")

type Stack interface {
	IsEmpty() bool
	Push(int)
	Pop() (int, error)
	Peek() (int, error)
}

func (ss SliceStack) String() string {
	return fmt.Sprint(ss.stack)
}

func (ss *SliceStack) IsEmpty() bool {
	return len(ss.stack) == 0
}

func (ss *SliceStack) Push(item int) {
	ss.stack = append(ss.stack, item)
}

func (ss *SliceStack) Pop() (int, error) {
	stackLen := len(ss.stack)
	if stackLen == 0 {
		return 0, ErrEmptyStack
	}

	top := ss.stack[stackLen-1]
	ss.stack = ss.stack[:stackLen-1]
	return top, nil
}

func (ss *SliceStack) Peek() (int, error) {
	stackLen := len(ss.stack)
	if stackLen == 0 {
		return 0, ErrEmptyStack
	}

	top := ss.stack[stackLen-1]
	return top, nil
}
