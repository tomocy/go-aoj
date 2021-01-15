package main

import "fmt"

func main() {
	var w, h int
	fmt.Scan(&w, &h)
	fmt.Println(solve(w, h))
}

func solve(w, h int) (int, int) {
	return w * h, (w + h) * 2
}
