package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	solve(os.Stdout, os.Stdin)
}

func solve(dst io.Writer, src io.Reader) {
	scanner := bufio.NewScanner(src)
	scanner.Scan()
	s := scanner.Text()

	b := new(strings.Builder)

	diff := 'a' - 'A'

	for _, r := range s {
		if 'A' <= r && r <= 'Z' {
			fmt.Fprint(b, string(r+diff))
			continue
		}

		if 'a' <= r && r <= 'z' {
			fmt.Fprint(b, string(r-diff))
			continue
		}

		fmt.Fprint(b, string(r))
	}

	fmt.Fprintln(dst, b)
}
