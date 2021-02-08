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
4
1 2
1 3
1 4
1 5
`,
			expected: `3
5
2
4
`,
		},
		{
			input: `1 2 3 4 5 6
4
6 2
6 3
6 4
6 5
`,
			expected: `4
2
5
3
`,
		},
		{
			input: `1 2 3 4 5 6
4
4 2
4 1
4 5
4 6
`,
			expected: `1
5
6
2
`,
		},
		{
			input: `1 2 3 4 5 6
4
3 1
3 2
3 6
3 5
`,
			expected: `2
6
5
1
`,
		},
		{
			input: `1 2 3 4 5 6
3
6 5
1 3
3 2
`,
			expected: `3
5
6
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
