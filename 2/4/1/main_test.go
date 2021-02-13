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
1 2 3 4 5
3
3 4 1
`,
			expected: "3\n",
		},
		{
			input: `3
3 1 2
1
5
`,
			expected: "0\n",
		},
		{
			input: `5
1 1 2 2 3
2
1 2
`,
			expected: "2\n",
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
