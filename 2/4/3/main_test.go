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
insert AAA
insert AAC
find AAA
find CCC
insert CCC
find CCC
`,
			expected: `yes
no
yes
`,
		},
		{
			input: `13
insert AAA
insert AAC
insert AGA
insert AGG
insert TTT
find AAA
find CCC
find CCC
insert CCC
find CCC
insert T
find TTT
find T
`,
			expected: `yes
no
no
yes
yes
yes
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
