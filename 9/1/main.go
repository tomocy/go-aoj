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
	var target string
	fmt.Fscan(src, &target)
	target = strings.ToLower(target)

	var words []string
	for {
		var word string
		fmt.Fscan(src, &word)
		if word == "END_OF_TEXT" {
			break
		}

		words = append(words, word)
	}

	var x int
	for _, word := range words {
		if strings.ToLower(word) == target {
			x++
		}
	}

	fmt.Fprintln(dst, x)
}
