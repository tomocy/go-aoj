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

	var (
		sorted  bool
		swapped int
	)
	for !sorted {
		sorted = true
		for i := len(vals) - 1; i >= 1; i-- {
			if vals[i] >= vals[i-1] {
				continue
			}

			vals[i-1], vals[i] = vals[i], vals[i-1]
			sorted = false
			swapped++
		}
	}

	b := new(bytes.Buffer)

	for i, v := range vals {
		if i != 0 {
			fmt.Fprint(b, " ")
		}
		fmt.Fprint(b, v)
	}
	fmt.Fprintln(b)

	fmt.Fprintln(b, swapped)

	fmt.Fprint(dst, b)
}
