package main

import "math"

func main() {}

func solve(n int) int {
	return pow(n, 3)
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
