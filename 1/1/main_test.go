package main

import (
	"bytes"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	expected := "Hello World"

	buf := new(bytes.Buffer)
	solve(buf)
	if buf.String() != expected {
		t.Errorf("got %v, but expected %v", buf, expected)
	}
}
