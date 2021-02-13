package main

import (
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

	s := make([]int, n)
	for i := range s {
		fmt.Fscan(src, &s[i])
	}

	fmt.Fscan(src, &n)

	t := make([]int, n)
	for i := range t {
		fmt.Fscan(src, &t[i])
	}

	var included int
	for _, vOfT := range t {
		for _, vOfS := range s {
			if vOfT == vOfS {
				included++
				break
			}
		}
	}

	fmt.Fprintln(dst, included)
}
