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
	var s, p string
	fmt.Fscan(src, &s)
	fmt.Fscan(src, &p)

	for i := 0; i < len(s); i++ {
		first := s[0]
		s = s[1:]
		s += string(first)

		if strings.Contains(s, p) {
			fmt.Fprintln(dst, "Yes")
			return
		}
	}

	fmt.Fprintln(dst, "No")
}
