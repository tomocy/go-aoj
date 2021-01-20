package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	fmt.Scan(&r)
	area, circum := solve(r)
	fmt.Printf("%f %f\n", area, circum)
}

func solve(r float64) (float64, float64) {
	return round(r*r*math.Pi, 6), round(2*r*math.Pi, 6)
}

func round(v float64, n int) float64 {
	p := math.Pow10(n)
	return math.Round(v*p) / p
}
