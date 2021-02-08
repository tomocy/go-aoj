package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		w              int
		h              int
		expectedArea   int
		expectedCircum int
	}{
		{
			w: 3, h: 5,
			expectedArea: 15, expectedCircum: 16,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(fmt.Sprintf("w,h:%v,%v", test.w, test.h), func(t *testing.T) {
			t.Parallel()

			actualArea, actualCircum := solve(test.w, test.h)
			if actualArea != test.expectedArea {
				t.Errorf("area: \nexpected %v\n     got %v", test.expectedArea, actualCircum)
			}
			if actualCircum != test.expectedCircum {
				t.Errorf("circum: \nexpected %v\n     got %v", test.expectedArea, actualCircum)
			}
		})
	}
}
