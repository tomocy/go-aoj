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
			r:        "3\n5\n11\n7\n8\n19\n0\n",
			expected: "Case 1: 3\nCase 2: 5\nCase 3: 11\nCase 4: 7\nCase 5: 8\nCase 6: 19\n",
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
