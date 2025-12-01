package basic_syntax

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// Benchmark Test
// When the benchmark code is executed, it measures how long it takes.
// After Loop() returns false, b.N contains the total number of iterations that ran.
func BenchmarkRepeat(b *testing.B) {
	// Setup code if needed
	for b.Loop() {
		//Code to measure
		Repeat("a", 5)
	}
	// Cleanup code if needed
}

func BenchmarkImprovedRepeat(b *testing.B) {
	// Setup code if needed
	for b.Loop() {
		//Code to measure
		ImprovedRepeat("a", 5)
	}
	// Cleanup code if needed
}

// Run in PowerShell or terminal
// go test -bench="."

// Testable Example
func ExampleRepeat() {
	result := Repeat("a", 5)
	fmt.Println(result)
	// Output: aaaaa
}

// run pkgsite -open ., to see the documentation with examples rendered
