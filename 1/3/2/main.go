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
	for i := 1; ; i++ {
		var x int
		_, err := fmt.Fscan(r, &x)
		if err != nil || x == 0 {
			return
		}

		fmt.Fprintf(w, "Case %d: %d\n", i, x)
	}
}
