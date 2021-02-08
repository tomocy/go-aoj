package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	solve(os.Stdout, os.Stdin)
}

func solve(dst io.Writer, src io.Reader) {
	var (
		x1, y1 float64
		x2, y2 float64
	)
	fmt.Fscan(src, &x1, &y1, &x2, &y2)

	d := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	d = float64(int(d*100000)) / 100000
	fmt.Fprintln(dst, d)
}
