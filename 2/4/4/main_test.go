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
			input: `5 3
8
1
7
3
9
`,
			expected: "10\n",
		},
		{
			input: `4 2
1
2
2
6
`,
			expected: "6\n",
		},
		{
			input: `10 4
5
3
2
1
1
6
3
4
5
7
`,
			expected: "11\n",
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
