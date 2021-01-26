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
		var n, x int
		fmt.Fscanln(src, &n, &x)

		if n == 0 && x == 0 {
			break
		}

		var ways int

		for i := 1; i <= n; i++ {
			for j := i + 1; j <= n; j++ {
				for k := j + 1; k <= n; k++ {
					if i+j+k == x {
						ways++
					}
				}
			}
		}

		fmt.Fprintln(b, ways)
	}

	fmt.Fprint(dst, b)
}
