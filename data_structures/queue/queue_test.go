package queue

import (
	"slices"
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("IsEmpty when queue is not empty", func(t *testing.T) {
		queue := SliceStringQueue{queue: []string{"1", "2", "3"}}

		got := queue.IsEmpty()
		want := false

		if got != want {
			t.Errorf("got %t but want %t", got, want)
		}
	})

	t.Run("IsEmpty when queue is empty", func(t *testing.T) {
		queue := SliceStringQueue{queue: []string{}}

		got := queue.IsEmpty()
		want := true

		if got != want {
			t.Errorf("got %t but want %t", got, want)
		}
	})

	t.Run("Enqueue to an empty queue", func(t *testing.T) {
		queue := SliceStringQueue{queue: []string{}}

		queue.Enqueue("5")
		want := []string{"5"}

		if !slices.Equal(queue.queue, want) {
			t.Errorf("got %v but want %v", queue.queue, want)
		}

	})

	t.Run("Successful dequeue", func(t *testing.T) {
		queue := SliceStringQueue{queue: []string{"1", "2", "3", "4"}}

		got, err := queue.Dequeue()
		want := "1"
		updatedQueue := []string{"2", "3", "4"}

		assertNoError(t, err)

		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}

		if !slices.Equal(queue.queue, updatedQueue) {
			t.Errorf("got %v but want %v", queue.queue, updatedQueue)
		}
	})

	t.Run("Dequeue error", func(t *testing.T) {
		queue := SliceStringQueue{queue: []string{}}

		_, err := queue.Dequeue()

		assertError(t, err, ErrEmptyQueue)
	})

	t.Run("Successful peek", func(t *testing.T) {
		queue := SliceStringQueue{queue: []string{"1", "2", "3", "4"}}

		got, err := queue.Peek()
		want := "1"
		updatedQueue := []string{"1", "2", "3", "4"}

		assertNoError(t, err)

		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}

		if !slices.Equal(queue.queue, updatedQueue) {
			t.Errorf("got %v but want %v", queue.queue, updatedQueue)
		}
	})

	t.Run("Peek error", func(t *testing.T) {
		queue := SliceStringQueue{queue: []string{}}

		_, err := queue.Peek()

		assertError(t, err, ErrEmptyQueue)
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
