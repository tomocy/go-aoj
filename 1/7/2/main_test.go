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
			input: `5 9
0 0
`,
			expected: "2\n",
		},
		{
			input: `3 0
3 1
3 2
3 3
3 4
3 5
3 6
3 7
30 36
30 27
30 28
30 11
30 12
30 23
30 24
30 30
30 26
40 40
40 40
40 80
40 50
40 30
40 20
40 10
40 60
40 70
40 90
50 55
50 65
50 75
50 85
50 95
50 10
99 120
99 130
99 140
99 150
99 160
100 88
100 56
100 26
100 36
100 116
100 120
100 100
100 186
100 286
100 208
0 0
`,
			expected: `0
0
0
0
0
0
1
0
87
48
52
5
7
33
37
61
44
114
114
133
168
61
24
4
190
178
75
223
278
300
288
243
4
1051
1134
1184
1201
1184
602
234
44
91
1015
1060
784
1027
16
705
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
