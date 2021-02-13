#!/bin/bash

set -eu

PROBLEM=${1:-""}

if [[ -z "$PROBLEM" ]]; then
  echo "no problem specified"
  exit 1
fi

PATH_MAIN_GO="$PROBLEM/main.go"
CONTENTS_MAIN_GO="package main

import (
	\"io\"
	\"os\"
)

func main() {
	solve(os.Stdout, os.Stdin)
}

func solve(dst io.Writer, src io.Reader) {}
"
PATH_MAIN_TEST_GO="$PROBLEM/main_test.go"
CONTENTS_MAIN_TEST_GO="package main

import (
	\"strings\"
	\"testing\"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    \`\`,
			expected: \`\`,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			actual := new(strings.Builder)
			solve(actual, strings.NewReader(test.input))
			if actual.String() != test.expected {
				t.Errorf(\"got %v, but expected %v\", actual, test.expected)
			}
		})
	}
}
"

mkdir -p "$PROBLEM"
echo "$CONTENTS_MAIN_GO" > "$PATH_MAIN_GO"
echo "$CONTENTS_MAIN_TEST_GO" > "$PATH_MAIN_TEST_GO"
