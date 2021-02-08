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

func solve(w io.Writer, r io.Reader) {
	buf := new(strings.Builder)

	for {
		var (
			a, b int
			op   byte
		)
		fmt.Fscanf(r, "%d %c %d\n", &a, &op, &b)

		if op == '?' {
			break
		}

		var x int
		switch op {
		case '+':
			x = a + b
		case '-':
			x = a - b
		case '*':
			x = a * b
		case '/':
			x = a / b
		}

		fmt.Fprintln(buf, x)
	}

	fmt.Fprint(w, buf)
}
