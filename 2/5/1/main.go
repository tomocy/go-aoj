package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	solve(os.Stdout, os.Stdin)
}

func solve(dst io.Writer, src io.Reader) {
	var n int
	fmt.Fscan(src, &n)

	vals := make([]int, n)
	for i := range vals {
		fmt.Fscan(src, &vals[i])
	}

	var q int
	fmt.Fscan(src, &q)

	targets := make([]int, q)
	for i := range targets {
		fmt.Fscan(src, &targets[i])
	}

	b := new(bytes.Buffer)

	for _, t := range targets {
		if canSum(vals, t) {
			fmt.Fprintln(b, "yes")
		} else {
			fmt.Fprintln(b, "no")
		}
	}

	fmt.Fprint(dst, b)
}

func canSum(vals []int, target int) bool {
	return canSumFrom(vals, target, 0)
}

func canSumFrom(vals []int, target, from int) bool {
	if target == 0 {
		return true
	}
	if from >= len(vals) {
		return false
	}

	return canSumFrom(vals, target, from+1) || canSumFrom(vals, target-vals[from], from+1)
}
