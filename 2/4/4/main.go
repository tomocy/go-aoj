package main

import (
	"fmt"
	"io"
	"os"
)

const (
	maxCargos = 100000
	maxWeight = 10000
	capAtMost = maxCargos * maxWeight
)

func main() {
	solve(os.Stdout, os.Stdin)
}

func solve(dst io.Writer, src io.Reader) {
	var cargos, trucks int
	fmt.Fscan(src, &cargos, &trucks)

	weights := make([]int, cargos)
	for i := range weights {
		fmt.Fscan(src, &weights[i])
	}

	var smallestCap int
	start, end := max(weights...), capAtMost
	for start <= end {
		cap := (start + end) / 2
		canLoad := canLoadAll(weights, trucks, cap)
		if !canLoadAll(weights, trucks, cap-1) && canLoad {
			smallestCap = cap
			break
		}

		if !canLoad {
			start = cap + 1
		} else {
			end = cap - 1
		}
	}

	fmt.Fprintln(dst, smallestCap)
}

func canLoadAll(weights []int, trucks, cap int) bool {
	var weightI int
	for i := 0; i < trucks; i++ {
		var sum int
		for ; weightI < len(weights); weightI++ {
			if sum+weights[weightI] > cap {
				break
			}
			sum += weights[weightI]
		}
	}

	return weightI == len(weights)
}

func max(vs ...int) int {
	var x int
	for i, v := range vs {
		if i == 0 || v > x {
			x = v
		}
	}
	return x
}
