package stack

import (
	"slices"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("IsEmpty_when_stack_is_not_empty", func(t *testing.T) {
		stack := SliceStack{stack: []int{1, 2, 3}}
		got := stack.IsEmpty()
		want := false

		if got != want {
			t.Errorf("got %t but want %t", got, want)
		}
	})

	t.Run("IsEmpty_when_stack_is_empty", func(t *testing.T) {
		stack := SliceStack{stack: []int{}}
		got := stack.IsEmpty()
		want := true

		if got != want {
			t.Errorf("got %t but want %t", got, want)
		}
	})

	t.Run("Push_to_empty_stack", func(t *testing.T) {
		stack := SliceStack{stack: []int{}}
		stack.Push(5)
		want := []int{5}

		if !slices.Equal(stack.stack, want) {
			t.Errorf("got %v but want %v", stack.stack, want)
		}
	})

	t.Run("Push_to_stack", func(t *testing.T) {
		stack := SliceStack{stack: []int{1, 2}}
		stack.Push(5)
		want := []int{1, 2, 5}

		if !slices.Equal(stack.stack, want) {
			t.Errorf("got %v but want %v", stack.stack, want)
		}
	})

	t.Run("Pop_error_on_empty_stack", func(t *testing.T) {
		stack := SliceStack{stack: []int{}}
		_, err := stack.Pop()
		assertError(t, err, ErrEmptyStack)
	})

	t.Run("Pop_stack_success", func(t *testing.T) {
		stack := SliceStack{stack: []int{1, 2, 3}}
		item, err := stack.Pop()
		assertNoError(t, err)

		want := 3
		expectedStack := []int{1, 2}

		if item != want {
			t.Errorf("got %d but want %d", item, want)
		}

		if !slices.Equal(stack.stack, expectedStack) {
			t.Errorf("got %v but want %v", stack.stack, expectedStack)
		}
	})

	t.Run("Peek_error_on_empty_stack", func(t *testing.T) {
		stack := SliceStack{stack: []int{}}
		_, err := stack.Peek()
		assertError(t, err, ErrEmptyStack)
	})

	t.Run("Peek_stack_success", func(t *testing.T) {
		stack := SliceStack{stack: []int{1, 2, 3}}
		item, err := stack.Peek()
		assertNoError(t, err)

		want := 3
		expectedStack := []int{1, 2, 3}

		if item != want {
			t.Errorf("got %d but want %d", item, want)
		}

		if !slices.Equal(stack.stack, expectedStack) {
			t.Errorf("got %v but want %v", stack.stack, expectedStack)
		}
	})

}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
