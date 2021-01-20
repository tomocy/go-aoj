package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	type expected struct {
		area   float64
		circum float64
	}
	tests := []struct {
		r        float64
		expected expected
	}{
		{
			r: 2,
			expected: expected{
				area:   12.566371,
				circum: 12.566371,
			},
		},
		{
			r: 3,
			expected: expected{
				area:   28.274334,
				circum: 18.849556,
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(strconv.FormatFloat(test.r, 'f', 8, 64), func(t *testing.T) {
			t.Parallel()

			actualArea, actualCircum := solve(test.r)
			if actualArea != test.expected.area {
				t.Errorf("got %v, but expected %v", actualArea, test.expected.area)
			}
			if actualCircum != test.expected.circum {
				t.Errorf("got %v, but expected %v", actualCircum, test.expected.circum)
			}
		})
	}
}
