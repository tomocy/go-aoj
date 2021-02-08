package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	solve(os.Stdout, os.Stdin)
}

func solve(w io.Writer, r io.Reader) {
	for {
		var a, b int
		_, err := fmt.Fscan(r, &a, &b)
		if err != nil || a == 0 && b == 0 {
			return
		}

		if a < b {
			fmt.Fprintln(w, a, b)
		} else {
			fmt.Fprintln(w, b, a)
		}
	}
}
