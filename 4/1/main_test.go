package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	type expected struct {
		q float64
		r int
	}
	tests := []struct {
		a, b     int
		expected expected
	}{
		{
			a: 3, b: 2,
			expected: expected{
				q: 1.50000,
				r: 1,
			},
		},
		{
			a: 58, b: 39,
			expected: expected{
				q: 1.4871794871794872,
				r: 19,
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(fmt.Sprint(test), func(t *testing.T) {
			t.Parallel()

			actualQ, actualR := solve(test.a, test.b)
			if actualQ != test.expected.q {
				t.Errorf("got %v, but expected %v", actualQ, test.expected.q)
			}
			if actualR != test.expected.r {
				t.Errorf("got %v, but expected %v", actualR, test.expected.r)
			}
		})
	}
}
