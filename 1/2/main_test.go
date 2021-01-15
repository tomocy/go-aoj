package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		n        int
		expected int
	}{
		{
			n: 2, expected: 8,
		},
		{
			n: 3, expected: 27,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(strconv.Itoa(test.n), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.n)
			if actual != test.expected {
				t.Errorf("\nexpected %v\n     got %v", test.expected, actual)
			}
		})
	}
}
