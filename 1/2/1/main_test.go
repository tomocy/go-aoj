package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		a, b     int
		expected string
	}{
		{
			a: 1, b: 2, expected: "a < b",
		},
		{
			a: 4, b: 3, expected: "a > b",
		},
		{
			a: 5, b: 5, expected: "a == b",
		},
	}

	for _, test := range tests {
		test := test

		t.Run(fmt.Sprintf("a,b:%v,%v", test.a, test.b), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.a, test.b)
			if actual != test.expected {
				t.Errorf("got %v, but expected %v", actual, test.expected)
			}
		})
	}
}
