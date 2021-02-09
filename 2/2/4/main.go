package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	solve(os.Stdout, os.Stdin)
}

func solve(dst io.Writer, src io.Reader) {
	input, err := ioutil.ReadAll(src)
	if err != nil {
		return
	}
	src = bytes.NewReader(input)

	var n int
	fmt.Fscan(src, &n)

	vals := make([]int, n)
	for i := range vals {
		fmt.Fscan(src, &vals[i])
	}

	var cnt int
	deltas := insertionDeltas(len(vals) / 2)
	for _, d := range deltas {
		cnt += insertionSort(vals, d)
	}

	b := new(bytes.Buffer)

	fmt.Fprintln(b, len(deltas))
	printVals(b, " ", deltas...)
	fmt.Fprintln(b, cnt)
	printVals(b, "\n", vals...)

	fmt.Fprint(dst, b)
}

func insertionSort(vals []int, d int) int {
	var cnt int

	for i := d; i < len(vals); i++ {
		v := vals[i]
		j := i - d
		for ; j >= 0 && vals[j] > v; j = j - d {
			vals[j+d] = vals[j]
			cnt++
		}
		vals[j+d] = v
	}

	return cnt
}

func insertionDeltas(max int) []int {
	var ds []int
	for i := max; i > 1; i /= 2 {
		ds = append(ds, i)
	}
	ds = append(ds, 1)

	return ds
}

func printVals(dst io.Writer, sep string, vs ...int) {
	for i, v := range vs {
		if i != 0 {
			fmt.Fprint(dst, sep)
		}
		fmt.Fprint(dst, v)
	}
	fmt.Fprintln(dst)
}
