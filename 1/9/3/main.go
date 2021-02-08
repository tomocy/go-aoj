package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	solve(os.Stdout, os.Stdin)
}

func solve(dst io.Writer, src io.Reader) {
	var (
		a, b int
	)

	var n int
	fmt.Fscan(src, &n)

	for i := 0; i < n; i++ {
		var aS, bS string
		fmt.Fscan(src, &aS, &bS)
		if aS == bS {
			a++
			b++
			continue
		}
		if aS < bS {
			b += 3
			continue
		}
		a += 3
	}

	fmt.Fprintf(dst, "%d %d\n", a, b)
}
