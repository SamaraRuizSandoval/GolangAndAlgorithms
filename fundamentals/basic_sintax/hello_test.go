package basic_syntax

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say Hello to a person", func(t *testing.T) {
		got := Hello("Alice", "")
		want := "Hello, Alice"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say Hello, World when name is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Carlos", "Spanish")
		want := "Hola, Carlos"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Marie", "French")
		want := "Bonjour, Marie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() //is needed to tell the test suite that this method is a helper.
	// By doing this, when it fails, the line number reported will be in our function call rather than inside our test helper.
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
