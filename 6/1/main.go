package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	solve(os.Stdout, os.Stdin)
}

func solve(dst io.Writer, src io.Reader) {
	var n int
	fmt.Fscan(src, &n)

	vs := make([]int, n)
	for i := range vs {
		var v int
		fmt.Fscan(src, &v)
		vs[i] = v
	}

	b := new(strings.Builder)

	for i := len(vs) - 1; i >= 0; i-- {
		fmt.Fprint(b, vs[i])
		if i != 0 {
			fmt.Fprint(b, " ")
		}
	}

	fmt.Fprintln(b)

	fmt.Fprint(dst, b)
}
