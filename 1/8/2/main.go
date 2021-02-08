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
	b := new(strings.Builder)

	for {
		var s string
		fmt.Fscan(src, &s)

		if s == "0" {
			break
		}

		fmt.Fprintln(b, sumDigits(s))
	}

	fmt.Fprint(dst, b)
}

func sumDigits(s string) int {
	var sum int

	for _, r := range s {
		sum += int(r - '0')
	}

	return sum
}
