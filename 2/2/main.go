package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(solve(a, b, c))
}

func solve(a, b, c int) string {
	if a < b && b < c {
		return "Yes"
	}
	return "No"
}
