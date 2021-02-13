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
	for _, v := range t {
		if binarySearch(s, v) {
			included++
		}
	}

	fmt.Fprintln(dst, included)
}

func binarySearch(vs []int, x int) bool {
	start, end := 0, len(vs)-1
	for start <= end {
		middle := (start + end) / 2
		if x == vs[middle] {
			return true
		}
		if x > vs[middle] {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return false
}
