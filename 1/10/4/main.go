package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
)

const roundRange = 10000

func main() {
	solve(os.Stdout, os.Stdin)
}

func solve(dst io.Writer, src io.Reader) {
	var n int
	fmt.Fscan(src, &n)

	var x []float64
	for i := 0; i < n; i++ {
		var v float64
		fmt.Fscan(src, &v)
		x = append(x, v)
	}

	var y []float64
	for i := 0; i < n; i++ {
		var v float64
		fmt.Fscan(src, &v)
		y = append(y, v)
	}

	b := new(bytes.Buffer)

	fmt.Fprintln(b, round(minkowski(x, y, 1)))
	fmt.Fprintln(b, round(minkowski(x, y, 2)))
	fmt.Fprintln(b, round(minkowski(x, y, 3)))
	fmt.Fprintln(b, round(minkowski(x, y, math.Inf(0))))

	fmt.Fprint(dst, b)
}

func minkowski(x, y []float64, p float64) float64 {
	if !math.IsInf(p, 0) {
		var sum float64
		for i := range x {
			sum += math.Pow(math.Abs(x[i]-y[i]), p)
		}

		return math.Pow(sum, 1/p)
	}

	var max float64
	for i := range x {
		v := math.Abs(x[i] - y[i])
		if i == 0 || v > max {
			max = v
		}
	}

	return max
}

func round(x float64) float64 {
	return math.Round(float64(int(x*roundRange))) / roundRange
}
