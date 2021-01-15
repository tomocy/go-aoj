package main

import "fmt"

func main() {
	var seconds int
	fmt.Scan(&seconds)
	fmt.Println(solve(seconds))
}

func solve(seconds int) string {
	h, rest := seconds/3600, seconds%3600
	m, s := rest/60, rest%60
	return fmt.Sprintf("%d:%d:%d", h, m, s)
}
