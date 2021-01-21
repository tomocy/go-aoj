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
		var h, w int
		fmt.Fscan(src, &h, &w)
		if h == 0 && w == 0 {
			break
		}

		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				if i%2 != 0 {
					if j%2 != 0 {
						fmt.Fprint(b, "#")
					} else {
						fmt.Fprint(b, ".")
					}
				} else {
					if j%2 != 0 {
						fmt.Fprint(b, ".")
					} else {
						fmt.Fprint(b, "#")
					}
				}
			}

			fmt.Fprintln(b)
		}

		fmt.Fprintln(b)
	}

	fmt.Fprint(dst, b)
}
