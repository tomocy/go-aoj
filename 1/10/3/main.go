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
	b := new(bytes.Buffer)

	for {
		var n int
		fmt.Fscan(src, &n)
		if n == 0 {
			break
		}

		var scores []int
		for i := 0; i < n; i++ {
			var score int
			fmt.Fscan(src, &score)
			scores = append(scores, score)
		}

		fmt.Fprintln(b, round(stdDeviation(scores...)))
	}

	fmt.Fprint(dst, b)
}

func stdDeviation(scores ...int) float64 {
	ave := average(scores...)

	var sum float64
	for _, score := range scores {
		sum += math.Pow(float64(score)-ave, 2)
	}

	return math.Sqrt(sum / float64(len(scores)))
}

func average(vs ...int) float64 {
	var sum float64
	for _, v := range vs {
		sum += float64(v)
	}

	return sum / float64(len(vs))
}

func round(x float64) float64 {
	return float64(int(x*roundRange)) / roundRange
}
