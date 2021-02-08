package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	solve(os.Stdout, 100)
}

func solve(w io.Writer, n int) {
	b := new(strings.Builder)
	for i := 0; i < n; i++ {
		fmt.Fprintln(b, "Hello World")
	}

	fmt.Fprint(w, b)
}
