package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected string
	}{
		{
			input: `5
H4 C9 S4 D2 C3`,
			expected: `D2 C3 H4 S4 C9
Stable
D2 C3 S4 H4 C9
Not stable
`,
		},
		{
			input: `2
S1 H1`,
			expected: `S1 H1
Stable
S1 H1
Stable
`,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			actual := new(strings.Builder)
			solve(actual, strings.NewReader(test.input))
			if actual.String() != test.expected {
				t.Errorf("got %v, but expected %v", actual, test.expected)
			}
		})
	}
}
