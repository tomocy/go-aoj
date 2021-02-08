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
	var n, m, l int
	fmt.Fscanln(src, &n, &m, &l)

	a, b := make([][]int, n), make([][]int, m)
	for i := range a {
		a[i] = make([]int, m)
	}
	for i := range b {
		b[i] = make([]int, l)
	}

	for i, row := range a {
		for j := range row {
			fmt.Fscan(src, &a[i][j])
		}
		fmt.Fscan(src)
	}

	for i, row := range b {
		for j := range row {
			fmt.Fscan(src, &b[i][j])
		}
		fmt.Fscan(src)
	}

	buf := new(strings.Builder)

	for i := 0; i < n; i++ {
		for j := 0; j < l; j++ {
			if j != 0 {
				fmt.Fprint(buf, " ")
			}
			fmt.Fprint(buf, product(a, b, m, i, j))
		}
		fmt.Fprintln(buf)
	}

	fmt.Fprint(dst, buf)
}

func product(a, b [][]int, m, i, j int) int {
	var x int

	for k := 0; k < m; k++ {
		x += a[i][k] * b[k][j]
	}

	return x
}
