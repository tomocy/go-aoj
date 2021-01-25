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
	var n, m int
	fmt.Fscanln(src, &n, &m)

	a, b := make([][]int, n), make([]int, m)
	for i := range a {
		a[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(src, &a[i][j])
		}

		fmt.Fscan(src)
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(src, &b[i])
	}

	buf := new(strings.Builder)

	for i := 0; i < n; i++ {
		fmt.Fprintln(buf, produce(a, b, i))
	}

	fmt.Fprint(dst, buf)
}

func produce(a [][]int, b []int, i int) int {
	var x int

	for j := range b {
		x += a[i][j] * b[j]
	}

	return x
}
