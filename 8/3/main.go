package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

func main() {
	solve(os.Stdout, os.Stdin)
}

func solve(dst io.Writer, src io.Reader) {
	read, err := ioutil.ReadAll(src)
	if err != nil {
		return
	}
	s := string(read)

	counts := make(map[rune]int)

	for _, r := range s {
		counts[unicode.ToLower(r)]++
	}

	b := new(strings.Builder)

	for i := 'a'; i <= 'z'; i++ {
		fmt.Fprintf(b, "%c : %d\n", i, counts[i])
	}

	fmt.Fprint(dst, b)
}
