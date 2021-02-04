package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected string
	}{
		{
			input: `5
70 80 100 90 20
3
80 80 80
0
`,
			expected: `27.8567
0
`,
		},
		{
			input: `5
70 80 100 90 20
3
80 80 80
1
100
1
50
1
0
10
0 0 0 0 0 0 0 0 0 0
10
1 34 44 63 30 1 9 53 57 57
10
20 12 52 44 6 9 94 31 67 70
0
`,
			expected: `27.8567
0
0
0
0
0
22.7396
28.4332
`,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			actual := new(strings.Builder)
			solve(actual, strings.NewReader(test.input))
			if actual.String() != test.expected {
				t.Errorf("got %v, but expected %v", actual, test.expected)
			}
		})
	}
}
