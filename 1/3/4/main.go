package main

import "fmt"

func main() {
	var begin, end, n int
	fmt.Scan(&begin, &end, &n)
	fmt.Println(solve(begin, end, n))
}

func solve(begin, end, n int) int {
	var x int

	for i := begin; i <= end; i++ {
		if n%i == 0 {
			x++
		}
	}

	return x
}
