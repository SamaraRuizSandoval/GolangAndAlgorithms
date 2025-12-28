package dgQueueProblems

import (
	"slices"
	"testing"
)

type GenerateBinaryNumbersCases struct {
	name  string
	input int
	want  []string
}

var tests = []GenerateBinaryNumbersCases{
	{name: "Example 1", input: 2, want: []string{"1", "10"}},
	{name: "Example 2", input: 3, want: []string{"1", "10", "11"}},
	{name: "Example 3", input: 5, want: []string{"1", "10", "11", "100", "101"}},
}

func TestGenerateBinaryNumbers(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateBinaryNumbers(tt.input)
			if !slices.Equal(got, tt.want) {
				t.Errorf("For input %q, expected %q, got %q", tt.input, tt.want, got)
			}
		})
	}
}

type Input struct {
	arr []int
	k   int
}

type PrintMaxCases struct {
	name  string
	input Input
	want  []int
}

var printMaxCasesTests = []PrintMaxCases{
	{
		name:  "Example 1",
		input: Input{arr: []int{1, 2, 3, 1, 4, 5, 2, 3, 6}, k: 3},
		want:  []int{3, 3, 4, 5, 5, 5, 6},
	},

	{
		name:  "Example 2",
		input: Input{arr: []int{8, 5, 10, 7, 9, 4, 15, 12, 90, 13}, k: 4},
		want:  []int{10, 10, 10, 15, 15, 90, 90},
	},

	{
		name:  "Example 3",
		input: Input{arr: []int{12, 1, 78, 90, 57}, k: 3},
		want:  []int{78, 90, 90},
	},

	{
		name:  "Example 4",
		input: Input{arr: []int{11, 10, 9, 8, 7}, k: 3},
		want:  []int{11, 10, 9},
	},
}

func TestPrintMaxCasesTests(t *testing.T) {
	for _, tt := range printMaxCasesTests {
		t.Run(tt.name, func(t *testing.T) {
			got := PrintMax(tt.input.arr, tt.input.k)
			if !slices.Equal(got, tt.want) {
				t.Errorf("For input %v, expected %v, got %v", tt.input, tt.want, got)
			}
		})
	}
}

type PalindromeCases struct {
	name  string
	input string
	want  bool
}

var palindromeTests = []PalindromeCases{
	{name: "Example 1", input: "madam", want: true},
	{name: "Example 2", input: "openai", want: false},
	{name: "Example 3", input: "A man a plan a canal Panama", want: true},
}

func TestPalindrome(t *testing.T) {
	for _, tt := range palindromeTests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckPalindrome(tt.input)
			if tt.want != got {
				t.Errorf("For input %v, expected %v, got %v", tt.input, tt.want, got)
			}
		})
	}
}
