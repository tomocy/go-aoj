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

	var n int
	fmt.Fscan(src, &n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			fmt.Fprint(b, " ", i)
			continue
		}

		for j := i; j > 0; j /= 10 {
			if j%10 == 3 {
				fmt.Fprint(b, " ", i)
				break
			}
		}
	}

	fmt.Fprintln(b)

	fmt.Fprint(dst, b)
}
