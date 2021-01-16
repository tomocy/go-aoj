package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		vs       []int
		expected []int
	}{
		{
			vs: []int{3, 8, 1}, expected: []int{1, 3, 8},
		},
		{
			vs: []int{10, 15, 10}, expected: []int{10, 10, 15},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(fmt.Sprint(test.vs), func(t *testing.T) {
			t.Parallel()

			actual := solve(test.vs)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("got %v, but expected %v", actual, test.expected)
			}
		})
	}
}
