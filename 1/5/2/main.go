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

func solve(dst io.Writer, r io.Reader) {
	b := new(strings.Builder)

	for {
		var h, w int
		fmt.Fscan(r, &h, &w)
		if h == 0 && w == 0 {
			break
		}

		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				if i == 0 || i == h-1 {
					fmt.Fprint(b, "#")
					continue
				}

				if j == 0 || j == w-1 {
					fmt.Fprint(b, "#")
					continue
				}

				fmt.Fprint(b, ".")
			}

			fmt.Fprintln(b)
		}

		fmt.Fprintln(b)
	}

	fmt.Fprint(dst, b)
}
