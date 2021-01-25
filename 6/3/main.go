package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	solve(os.Stdout, os.Stdin)
}

func solve(dst io.Writer, src io.Reader) {
	inDom := make([][][]int, 4)
	for i := range inDom {
		inDom[i] = make([][]int, 3)
		for j := range inDom[i] {
			inDom[i][j] = make([]int, 10)
		}
	}

	var n int
	fmt.Fscanln(src, &n)

	for i := 0; i < n; i++ {
		var dom, floor, room, people int
		fmt.Fscanln(src, &dom, &floor, &room, &people)

		inDom[dom-1][floor-1][room-1] += people
	}

	b := new(strings.Builder)

	for i, inFloor := range inDom {
		for _, inRoom := range inFloor {
			for _, people := range inRoom {
				fmt.Fprintf(b, " %d", people)
			}

			fmt.Fprintln(b)
		}

		if i != len(inDom)-1 {
			fmt.Fprintln(b, "####################")
		}
	}

	fmt.Fprint(dst, b)
}
