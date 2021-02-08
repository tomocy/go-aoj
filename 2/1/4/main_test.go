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
			input: `6
5
3
1
3
4
3`,
			expected: "3\n",
		},
		{
			input: `3
4
3
2`,
			expected: "-1\n",
		},
		{
			input: `4
4
1
2
1`,
			expected: "1\n",
		},
		{
			input: `7
5
4
3
2
4
3
1`,
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
