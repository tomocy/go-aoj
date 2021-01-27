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
			input: `123
55
1000
0
`,
			expected: `6
10
1
`,
		},
		{
			input: `123434134414
999999999999999999
234823492348923949234
1134842394
12349
219349
2394295922
2134324009823984932973957384823684234729917374928374
89237497289349179374923947926525620343
8412634873268747123
2987489327983
1000000000000000000000001
0
`,
			expected: `34
162
97
39
19
28
47
253
196
86
79
2
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
