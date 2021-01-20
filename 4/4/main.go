package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	solve(os.Stdout, os.Stdin)
}

func solve(w io.Writer, r io.Reader) {
	var n int
	fmt.Fscan(r, &n)

	vs := make([]int, n)
	for i := range vs {
		fmt.Fscan(r, &vs[i])
	}

	fmt.Fprintln(w, min(vs...), max(vs...), sum(vs...))
}

func min(vs ...int) int {
	var min int

	for i, v := range vs {
		if i == 0 || v < min {
			min = v
		}
	}

	return min
}

func max(vs ...int) int {
	var max int

	for i, v := range vs {
		if i == 0 || v > max {
			max = v
		}
	}

	return max
}

func sum(vs ...int) int {
	var sum int

	for _, v := range vs {
		sum += v
	}

	return sum
}
