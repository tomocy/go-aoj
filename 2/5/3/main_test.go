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
			input: "1",
			expected: `0.000000 0.000000
33.333333 0.000000
50.000000 28.867513
66.666667 0.000000
100.000000 0.000000
`,
		},
		{
			input: "2",
			expected: `0.000000 0.000000
11.111111 0.000000
16.666667 9.622504
22.222222 0.000000
33.333333 0.000000
38.888889 9.622504
33.333333 19.245009
44.444444 19.245009
50.000000 28.867513
55.555556 19.245009
66.666667 19.245009
61.111111 9.622504
66.666667 0.000000
77.777778 0.000000
83.333333 9.622504
88.888889 0.000000
100.000000 0.000000
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
