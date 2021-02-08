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

	b := new(bytes.Buffer)

	for i, v := range vals {
		j := i - 1
		for j >= 0 && vals[j] > v {
			vals[j+1] = vals[j]
			j--
		}
		vals[j+1] = v

		printlnVals(b, vals...)
	}

	fmt.Fprint(dst, b)
}

func printlnVals(dst io.Writer, vs ...int) {
	for i, v := range vs {
		if i != 0 {
			fmt.Fprint(dst, " ")
		}
		fmt.Fprint(dst, v)
	}
	fmt.Fprintln(dst)
}
