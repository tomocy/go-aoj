package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		seconds  int
		expected string
	}{
		{
			seconds: 46979, expected: "13:2:59",
		},
	}

	for _, test := range tests {
		test := test

		t.Run(strconv.Itoa(test.seconds), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.seconds)
			if actual != test.expected {
				t.Errorf("got %v, but expected %v", actual, test.expected)
			}
		})
	}
}
