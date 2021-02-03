package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	solve(os.Stdout, os.Stdin)
}

func solve(dst io.Writer, src io.Reader) {
	b := new(bytes.Buffer)

	for {
		var original string
		fmt.Fscan(src, &original)
		if original == "-" {
			break
		}

		var shuffleTimes int
		fmt.Fscan(src, &shuffleTimes)

		for i := 0; i < shuffleTimes; i++ {
			var h int
			fmt.Fscan(src, &h)

			toTop := original[:h]
			original = original[h:] + toTop
		}

		fmt.Fprintln(b, original)
	}

	fmt.Fprint(dst, b)
}
