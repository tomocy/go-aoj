package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

var (
	pointStart = point{
		x: 0, y: 0,
	}
	pointEnd = point{
		x: 100, y: 0,
	}
)

func main() {
	solve(os.Stdout, os.Stdin)
}

func solve(dst io.Writer, src io.Reader) {
	var n int
	fmt.Fscan(src, &n)

	points := []point{
		pointStart,
	}
	koshCurve(&points, n, pointStart, pointEnd)
	points = append(points, pointEnd)

	b := bufio.NewWriter(dst)

	for _, p := range points {
		fmt.Fprintf(b, "%f %f\n", p.x, p.y)
	}

	b.Flush()
}

func koshCurve(dst *[]point, n int, a, b point) {
	if n == 0 {
		return
	}

	s, t := divideInternally(a, b, 1, 2), divideInternally(a, b, 2, 1)
	u := point{
		x: (t.x-s.x)*math.Cos(radian(60)) - (t.y-s.y)*math.Sin(radian(60)) + s.x,
		y: (t.x-s.x)*math.Sin(radian(60)) + (t.y-s.y)*math.Cos(radian(60)) + s.y,
	}

	koshCurve(dst, n-1, a, s)
	*dst = append(*dst, s)
	koshCurve(dst, n-1, s, u)
	*dst = append(*dst, u)
	koshCurve(dst, n-1, u, t)
	*dst = append(*dst, t)
	koshCurve(dst, n-1, t, b)
}

func divideInternally(a, b point, m, n int) point {
	return point{
		x: (float64(n)*a.x + float64(m)*b.x) / (float64(m) + float64(n)),
		y: (float64(n)*a.y + float64(m)*b.y) / (float64(m) + float64(n)),
	}
}

func radian(degree float64) float64 {
	return degree * math.Pi / 180
}

type point struct {
	x, y float64
}
