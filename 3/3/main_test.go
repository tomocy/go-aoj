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
			r:        "3 2\n2 2\n5 3\n0 0\n",
			expected: "2 3\n2 2\n3 5\n",
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
