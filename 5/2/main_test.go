package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		r        string
		expected string
	}{
		{
			r: `3 4
5 6
3 3
0 0
`,
			expected: `####
#..#
####

######
#....#
#....#
#....#
######

###
#.#
###

`,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.r, func(t *testing.T) {
			t.Parallel()

			actual := new(strings.Builder)
			solve(actual, strings.NewReader(test.r))
			if actual.String() != test.expected {
				t.Errorf("got %v, but expected %v", actual, test.expected)
			}
		})
	}
}
