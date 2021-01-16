package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	solved := solve([]int{a, b, c})
	for i, v := range solved {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}

func solve(vs []int) []int {
	if len(vs) < 2 {
		return vs
	}

	pivot, rest := vs[0], vs[1:]
	less, greater := make([]int, 0, len(vs)), make([]int, 0, len(rest))
	for _, v := range rest {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}
	less, greater = solve(less), solve(greater)

	less = append(less, pivot)
	less = append(less, greater...)
	return less
}
