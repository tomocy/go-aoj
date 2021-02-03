package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	solve(os.Stdout, os.Stdin)
}

func solve(dst io.Writer, src io.Reader) {
	var s string
	fmt.Fscan(src, &s)

	var n int
	fmt.Fscan(src, &n)

	buf := new(bytes.Buffer)

	for i := 0; i < n; i++ {
		var (
			cmd  string
			a, b int
		)
		fmt.Fscan(src, &cmd, &a, &b)

		switch cmd {
		case "print":
			fmt.Fprintln(buf, s[a:b+1])
			continue
		case "reverse":
			left, right := s[:a], s[b+1:]
			var reversed []byte
			for i := b; i >= a; i-- {
				reversed = append(reversed, s[i])
			}
			s = left + string(reversed) + right
			continue
		case "replace":
			var p string
			fmt.Fscan(src, &p)
			s = strings.Replace(s, s[a:b+1], p, 1)
			continue
		}
	}

	fmt.Fprint(dst, buf)
}
