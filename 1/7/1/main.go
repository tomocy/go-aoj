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
	b := new(strings.Builder)

	for {
		var midterm, final, makeup int
		fmt.Fscanln(src, &midterm, &final, &makeup)

		if midterm == -1 && final == -1 && makeup == -1 {
			break
		}

		if midterm == -1 || final == -1 {
			fmt.Fprintln(b, "F")
			continue
		}

		if midterm+final >= 80 {
			fmt.Fprintln(b, "A")
			continue
		}

		if midterm+final >= 65 {
			fmt.Fprintln(b, "B")
			continue
		}

		if midterm+final >= 50 {
			fmt.Fprintln(b, "C")
			continue
		}

		if midterm+final >= 30 && makeup >= 50 {
			fmt.Fprintln(b, "C")
			continue
		}

		if midterm+final >= 30 {
			fmt.Fprintln(b, "D")
			continue
		}

		fmt.Fprintln(b, "F")
	}

	fmt.Fprint(dst, b)
}
