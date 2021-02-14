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
1 5 7 10 21
4
2 4 17 8
`,
			expected: `no
no
yes
yes
`,
		},
		{
			input: `5
1 5 7 10 21
10
2 4 6 15 17 8 22 21 100 35
`,
			expected: `no
no
yes
yes
yes
yes
yes
yes
no
no
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
