package main

import (
	"bufio"
	"bytes"
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

	sort(vals)

	compared := sort(vals)

	b := new(bytes.Buffer)

	printlnVals(b, vals...)
	fmt.Fprintln(b, compared)

	fmt.Fprint(dst, b)
}

func sort(vals []int) int {
	return mergeSort(vals, 0, len(vals))
}

func mergeSort(vals []int, start, end int) int {
	if end-start <= 1 {
		return 0
	}

	var compared int
	mid := (start + end) / 2
	compared += mergeSort(vals, start, mid)
	compared += mergeSort(vals, mid, end)
	compared += merge(vals, start, mid, end)

	return compared
}

func merge(vals []int, start, mid, end int) int {
	leftN, rightN := mid-start, end-mid
	left, right := make([]int, leftN+1), make([]int, rightN+1)
	for i := 0; i < leftN; i++ {
		left[i] = vals[start+i]
	}
	left[leftN] = math.MaxInt64
	for i := 0; i < rightN; i++ {
		right[i] = vals[mid+i]
	}
	right[rightN] = math.MaxInt64

	var compared int
	for i, leftI, rightI := start, 0, 0; i < end; i++ {
		if left[leftI] < right[rightI] {
			vals[i] = left[leftI]
			leftI++
		} else {
			vals[i] = right[rightI]
			rightI++
		}

		compared++
	}

	return compared
}

func printlnVals(dst io.Writer, vals ...int) {
	for i, v := range vals {
		if i != 0 {
			fmt.Fprint(dst, " ")
		}
		fmt.Fprint(dst, v)
	}
	fmt.Fprintln(dst)
}
