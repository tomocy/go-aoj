package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(solve(a, b))
}

func solve(a, b int) string {
	if a == b {
		return "a == b"
	}
	if a > b {
		return "a > b"
	}
	return "a < b"
}
