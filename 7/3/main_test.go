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
			input: `4 5
1 1 3 4 5
2 2 2 4 5
3 3 0 1 1
2 3 4 4 6
`,
			expected: `1 1 3 4 5 14
2 2 2 4 5 15
3 3 0 1 1 8
2 3 4 4 6 19
8 9 9 13 17 56
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
