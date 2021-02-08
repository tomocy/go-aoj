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
			input: `1 2 3 4 5 6
6 2 4 3 5 1
`,
			expected: "Yes\n",
		},
		{
			input: `1 2 3 4 5 6
6 5 4 3 2 1
`,
			expected: "No\n",
		},
		{
			input: `1 1 1 1 1 6
1 6 1 1 1 1
`,
			expected: "Yes\n",
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
