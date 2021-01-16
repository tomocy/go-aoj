package main

import "fmt"

func main() {
	var w, h, x, y, r int
	fmt.Scan(&w, &h, &x, &y, &r)
	fmt.Println(solve(w, h, x, y, r))
}

func solve(w, h, x, y, r int) string {
	top, bottom, left, right := y+r, y-r, x-r, x+r
	if 0 <= top && top <= h && 0 <= bottom && bottom <= h &&
		0 <= left && left <= w && 0 <= right && right <= w {
		return "Yes"
	}
	return "No"
}
