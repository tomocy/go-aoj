package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var markIndexes = map[string]int{
	"S": 0,
	"H": 1,
	"C": 2,
	"D": 3,
}

func main() {
	solve(os.Stdout, os.Stdin)
}

func solve(dst io.Writer, src io.Reader) {
	cards := make([][]bool, 4)
	for i := range cards {
		cards[i] = make([]bool, 13)
	}

	var n int
	fmt.Fscanln(src, &n)

	for i := 0; i < n; i++ {
		var (
			mark string
			rank int
		)
		fmt.Fscanln(src, &mark, &rank)

		cards[markIndexes[mark]][rank-1] = true
	}

	b := new(strings.Builder)

	for i, marked := range cards {
		for rank, inHand := range marked {
			if inHand {
				continue
			}

			fmt.Fprintf(b, "%s %d\n", markOfIndex(i), rank+1)
		}
	}

	fmt.Fprint(dst, b)
}

func markOfIndex(index int) string {
	for m, i := range markIndexes {
		if i == index {
			return m
		}
	}

	return ""
}
