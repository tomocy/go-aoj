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

	var swapped int
	for i := range vals {
		minIdx := i
		for j := i; j < len(vals); j++ {
			if vals[j] < vals[minIdx] {
				minIdx = j
			}
		}

		if vals[i] != vals[minIdx] {
			vals[i], vals[minIdx] = vals[minIdx], vals[i]
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
