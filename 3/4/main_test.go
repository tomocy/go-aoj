package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		begin, end int
		n          int
		expected   int
	}{
		{
			begin: 5, end: 14,
			n:        80,
			expected: 3,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(fmt.Sprintf("begin,end:%v,%v", test.begin, test.end), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.begin, test.end, test.n)
			if actual != test.expected {
				t.Errorf("got %v, but expected %v", actual, test.expected)
			}
		})
	}
}
