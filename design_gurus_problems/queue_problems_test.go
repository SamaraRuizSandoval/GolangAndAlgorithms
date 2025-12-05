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
