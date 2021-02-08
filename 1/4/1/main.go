package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	q, r := solve(a, b)
	fmt.Printf("%d %d %f\n", int(q), r, q)
}

func solve(a, b int) (float64, int) {
	return float64(a) / float64(b), a % b
}
