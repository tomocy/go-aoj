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
	var n int
	fmt.Fscan(src, &n)

	vals := make([]int, n)
	for i := range vals {
		fmt.Fscan(src, &vals[i])
	}

	var count int
	for _, v := range vals {
		if isPrime(v) {
			count++
		}
	}

	fmt.Fprintln(dst, count)
}

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	if x == 2 {
		return true
	}
	if x%2 == 0 {
		return false
	}

	sqrtX := int(math.Sqrt(float64(x)))

	for i := 3; i <= sqrtX; i += 2 {
		if x%i == 0 {
			return false
		}
	}

	return true
}
