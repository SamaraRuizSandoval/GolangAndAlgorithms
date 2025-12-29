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
		want := "5 → 10 → 20 → 30 → 40 → "

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

	t.Run("Delete the front node", func(t *testing.T) {
		buffer := bytes.Buffer{}
		err := list.DeleteAtPosition(0)
		list.Traverse(&buffer)

		got := buffer.String()
		want := "10 → 20 → 30 → 40 → "

		assertNoError(t, err)
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("Delete empty list", func(t *testing.T) {
		emptyList := SinglyLinkedList{}
		err := emptyList.DeleteAtPosition(0)

		assertError(t, err, ErrInvalidPosition)

	})

	t.Run("Delete a node in the middle", func(t *testing.T) {
		buffer := bytes.Buffer{}
		err := list.DeleteAtPosition(1)
		list.Traverse(&buffer)

		got := buffer.String()
		want := "10 → 30 → 40 → "

		assertNoError(t, err)
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("Delete a node in the tail", func(t *testing.T) {
		buffer := bytes.Buffer{}
		err := list.DeleteAtPosition(2)
		list.Traverse(&buffer)

		got := buffer.String()
		want := "10 → 30 → "

		assertNoError(t, err)
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("Delete a node in an invalid position", func(t *testing.T) {
		err := list.DeleteAtPosition(7)

		assertError(t, err, ErrInvalidPosition)

	})
}

func TestDoubleLinkedList(t *testing.T) {
	list := DoubleLinkedList{Head: &DoubleListNode{Val: 10, Prev: nil, Next: nil}}
	list.Head.Next = &DoubleListNode{Val: 20, Prev: list.Head, Next: nil}
	list.Head.Next.Next = &DoubleListNode{Val: 30, Prev: list.Head.Next, Next: nil}
	t.Run("Traverse List Forward", func(t *testing.T) {
		buffer := bytes.Buffer{}
		list.TraverseForward(&buffer)

		got := buffer.String()
		want := "10 ⇄ 20 ⇄ 30 ⇄ "

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("Traverse List Backward", func(t *testing.T) {
		buffer := bytes.Buffer{}
		list.TraverseBackward(&buffer)

		got := buffer.String()
		want := "30 ⇄ 20 ⇄ 10 ⇄ "

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("Insert value in the middle", func(t *testing.T) {
		buffer := bytes.Buffer{}
		list.InsertAtPosition(25, 2)
		list.TraverseForward(&buffer)

		got := buffer.String()
		want := "10 ⇄ 20 ⇄ 25 ⇄ 30 ⇄ "

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("Insert at the front", func(t *testing.T) {
		buffer := bytes.Buffer{}
		list.InsertAtPosition(5, 0)
		list.TraverseForward(&buffer)

		got := buffer.String()
		want := "5 ⇄ 10 ⇄ 20 ⇄ 25 ⇄ 30 ⇄ "

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("Insert at the front emptyList", func(t *testing.T) {
		newList := DoubleLinkedList{Head: nil}
		buffer := bytes.Buffer{}
		err := newList.InsertAtPosition(25, 0)
		newList.TraverseForward(&buffer)

		got := buffer.String()
		want := "25 ⇄ "

		assertNoError(t, err)
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("Delete the front node", func(t *testing.T) {
		buffer := bytes.Buffer{}
		err := list.DeleteAtPosition(0)
		list.TraverseForward(&buffer)

		got := buffer.String()
		want := "10 ⇄ 20 ⇄ 25 ⇄ 30 ⇄ "

		assertNoError(t, err)
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("Delete empty list", func(t *testing.T) {
		emptyList := DoubleLinkedList{}
		err := emptyList.DeleteAtPosition(0)

		assertError(t, err, ErrInvalidPosition)

	})

	t.Run("Delete list of one element", func(t *testing.T) {
		emptyList := DoubleLinkedList{Head: &DoubleListNode{Val: 10, Prev: nil, Next: nil}}
		err := emptyList.DeleteAtPosition(0)

		assertNoError(t, err)

	})

	t.Run("Delete a node in the middle", func(t *testing.T) {
		buffer := bytes.Buffer{}
		err := list.DeleteAtPosition(1)
		list.TraverseForward(&buffer)

		got := buffer.String()
		want := "10 ⇄ 25 ⇄ 30 ⇄ "

		assertNoError(t, err)
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("Delete a node in the tail", func(t *testing.T) {
		buffer := bytes.Buffer{}
		err := list.DeleteAtPosition(2)
		list.TraverseForward(&buffer)

		got := buffer.String()
		want := "10 ⇄ 25 ⇄ "

		assertNoError(t, err)
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("Delete a node in an invalid position", func(t *testing.T) {
		err := list.DeleteAtPosition(7)

		assertError(t, err, ErrInvalidPosition)

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
