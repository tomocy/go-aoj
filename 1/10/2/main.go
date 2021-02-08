package main

import (
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
	var (
		a, b float64
		c    int
	)
	fmt.Fscan(src, &a, &b, &c)

	r := radian(float64(c))

	buf := new(bytes.Buffer)

	s := a * b * math.Sin(r) / 2
	fmt.Fprintln(buf, s)

	l := a + b + math.Sqrt(math.Pow(a, 2)+math.Pow(b, 2)-2*a*b*math.Cos(r))
	fmt.Fprintln(buf, l)

	h := b * math.Sin(r)
	fmt.Fprintln(buf, h)

	fmt.Fprint(dst, buf)
}

func radian(theta float64) float64 {
	return theta * math.Pi / 180
}
