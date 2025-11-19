package basic_syntax

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// Benchmark Test
//When the benchmark code is executed, it measures how long it takes.
// After Loop() returns false, b.N contains the total number of iterations that ran.
func BenchmarkRepeat(b *testing.B) {
	// Setup code if needed
	for b.Loop() {
		//Code to measure
		Repeat("a")
	}
	// Cleanup code if needed
}

func BenchmarkImprovedRepeat(b *testing.B) {
	// Setup code if needed
	for b.Loop() {
		//Code to measure
		ImprovedRepeat("a")
	}
	// Cleanup code if needed
}
