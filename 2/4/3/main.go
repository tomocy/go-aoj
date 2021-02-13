package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	solve(os.Stdout, os.Stdin)
}

func solve(dst io.Writer, src io.Reader) {
	src = bufio.NewReader(src)

	var n int
	fmt.Fscan(src, &n)

	b := new(bytes.Buffer)

	dict := make(map[string]struct{})
	for i := 0; i < n; i++ {
		var verb, obj string
		fmt.Fscan(src, &verb, &obj)
		switch verb {
		case "insert":
			dict[obj] = struct{}{}
		case "find":
			_, ok := dict[obj]
			if ok {
				fmt.Fprintln(b, "yes")
			} else {
				fmt.Fprintln(b, "no")
			}
		}
	}

	fmt.Fprint(dst, b)
}
