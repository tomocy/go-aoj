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
			input: `abcde
3
replace 1 3 xyz
reverse 0 2
print 1 4
`,
			expected: "xaze\n",
		},
		{
			input: `xyz
3
print 0 2
replace 0 2 abc
print 0 2
`,
			expected: `xyz
abc
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
