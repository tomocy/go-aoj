package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	solve(os.Stdout, os.Stdin)
}

func solve(dst io.Writer, src io.Reader) {
	src = bufio.NewReader(src)

	var n int
	fmt.Fscan(src, &n)

	vals := make([]int, n)
	for i := range vals {
		fmt.Fscan(src, &vals[i])
	}

	inv := mergeSort(vals, 0, len(vals))

	fmt.Fprintln(dst, inv)
}

func mergeSort(vals []int, start, end int) int {
	if end-start <= 1 {
		return 0
	}

	mid := (start + end) / 2
	var inv int
	inv += mergeSort(vals, start, mid)
	inv += mergeSort(vals, mid, end)
	inv += merge(vals, start, mid, end)

	return inv
}

func merge(vals []int, start, mid, end int) int {
	leftN, rightN := mid-start, end-mid
	left, right := make([]int, leftN+1), make([]int, rightN+1)
	copy(left, vals[start:mid])
	left[leftN] = math.MaxInt64
	copy(right, vals[mid:end])
	right[rightN] = math.MaxInt64

	var inv int

	for i, leftI, rightI := start, 0, 0; i < end; i++ {
		if left[leftI] < right[rightI] {
			vals[i] = left[leftI]
			leftI++
		} else {
			vals[i] = right[rightI]
			rightI++
			inv += leftN - leftI
		}
	}

	return inv
}
