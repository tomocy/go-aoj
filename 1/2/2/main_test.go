package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		a, b, c  int
		expected string
	}{
		{
			a: 1, b: 3, c: 8, expected: "Yes",
		},
		{
			a: 3, b: 8, c: 1, expected: "No",
		},
	}

	for _, test := range tests {
		test := test

		t.Run(fmt.Sprintf("a,b,c,:%v,%v,%v", test.a, test.b, test.c), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.a, test.b, test.c)
			if actual != test.expected {
				t.Errorf("got %v, but expected %v", actual, test.expected)
			}
		})
	}
}
