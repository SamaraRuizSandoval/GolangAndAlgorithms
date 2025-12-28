package linkedlist

import (
	"bytes"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	list := SinglyLinkedList{Head: &ListNode{Val: 10, Next: &ListNode{Val: 20, Next: &ListNode{Val: 40, Next: nil}}}}
	t.Run("TraverseList", func(t *testing.T) {
		buffer := bytes.Buffer{}
		list.Traverse(&buffer)

		got := buffer.String()
		want := "10 → 20 → 40 → "

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("Insert value in the middle", func(t *testing.T) {
		buffer := bytes.Buffer{}
		list.InsertAtPosition(30, 2)
		list.Traverse(&buffer)

		got := buffer.String()
		want := "10 → 20 → 30 → 40 → "

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("Insert at the front", func(t *testing.T) {
		buffer := bytes.Buffer{}
		list.InsertAtPosition(5, 0)
		list.Traverse(&buffer)

		got := buffer.String()
		want := "5 → 10 → 20 → 40 → "

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("Insert at the front emptyList", func(t *testing.T) {
		newList := SinglyLinkedList{}
		buffer := bytes.Buffer{}
		err := newList.InsertAtPosition(25, 0)
		newList.Traverse(&buffer)

		got := buffer.String()
		want := "25 → "

		assertNoError(t, err)
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("Insert at the front emptyList", func(t *testing.T) {
		newList := SinglyLinkedList{}
		buffer := bytes.Buffer{}
		err := newList.InsertAtPosition(25, 0)
		newList.Traverse(&buffer)

		got := buffer.String()
		want := "25 → "

		assertNoError(t, err)
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("Insert in an invalid position", func(t *testing.T) {
		err := list.InsertAtPosition(25, 8)

		assertError(t, err, ErrInvalidPosition)
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
