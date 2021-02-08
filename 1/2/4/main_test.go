package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		w, h     int
		x, y, r  int
		expected string
	}{
		{
			w: 5, h: 4,
			x: 2, y: 2, r: 1,
			expected: "Yes",
		},
		{
			w: 5, h: 4,
			x: 2, y: 4, r: 1,
			expected: "No",
		},
	}

	for _, test := range tests {
		test := test

		t.Run(fmt.Sprint(test), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.w, test.h, test.x, test.y, test.r)
			if actual != test.expected {
				t.Errorf("got %v, but expected %v", actual, test.expected)
			}
		})
	}
}
