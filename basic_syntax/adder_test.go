package basic_syntax

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}

}

// Testable Example
func ExampleAdd() {
	result := Add(3, 5)
	fmt.Println(result)
	// Output: 8
}
