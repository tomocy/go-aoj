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
			input: `1 2 3 4 5 6

		`,
			expected: "1\n",
		},
		{
			input: `1 2 3 4 5 6
N
		`,
			expected: "2\n",
		},
		{
			input: `1 2 3 4 5 6
NN
		`,
			expected: "6\n",
		},
		{
			input: `1 2 3 4 5 6
NNN
		`,
			expected: "5\n",
		},
		{
			input: `1 2 3 4 5 6
NNNN
		`,
			expected: "1\n",
		},
		{
			input: `1 2 3 4 5 6
S
		`,
			expected: "5\n",
		},
		{
			input: `1 2 3 4 5 6
SS
		`,
			expected: "6\n",
		},
		{
			input: `1 2 3 4 5 6
SSS
		`,
			expected: "2\n",
		},
		{
			input: `1 2 3 4 5 6
SSSS
		`,
			expected: "1\n",
		},
		{
			input: `1 2 3 4 5 6
W
		`,
			expected: "3\n",
		},
		{
			input: `1 2 3 4 5 6
WW
		`,
			expected: "6\n",
		},
		{
			input: `1 2 3 4 5 6
WWW
		`,
			expected: "4\n",
		},
		{
			input: `1 2 3 4 5 6
WWWW
		`,
			expected: "1\n",
		},
		{
			input: `1 2 3 4 5 6
E
		`,
			expected: "4\n",
		},
		{
			input: `1 2 3 4 5 6
EE
		`,
			expected: "6\n",
		},
		{
			input: `1 2 3 4 5 6
EEE
		`,
			expected: "3\n",
		},
		{
			input: `1 2 3 4 5 6
EEEE
		`,
			expected: "1\n",
		},
		{
			input: `1 2 4 8 16 32
SE
		`,
			expected: "8\n",
		},
		{
			input: `1 2 4 8 16 32
EESWN
		`,
			expected: "32\n",
		},
		{
			input: `6 5 4 3 2 1
SSSSEEENNNWWSEWNNWNWESSE
		`,
			expected: "4\n",
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
