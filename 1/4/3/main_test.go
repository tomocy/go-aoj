package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		r        string
		expected string
	}{
		{
			r: `1 + 2
56 - 18
13 * 2
100 / 10
27 + 81
0 ? 0
`,
			expected: `3
38
26
10
108
`,
		},
		{
			r: `20000 + 20000
20000 / 20000
20000 * 20000
20000 - 20000
0 + 20000
0 / 20000
0 * 20000
0 - 20000
10009 * 997
10009 - 997
10009 + 997
10009 / 997
8 * 8
8 / 8
8 - 8
8 + 8
100 ? 100
`,
			expected: `40000
1
400000000
0
20000
0
0
-20000
9978973
9012
11006
10
64
1
0
16
`,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.r, func(t *testing.T) {
			t.Parallel()

			actual := new(strings.Builder)
			solve(actual, strings.NewReader(test.r))
			if actual.String() != test.expected {
				t.Errorf("got %v, but expected %v", actual, test.expected)
			}
		})
	}
}
