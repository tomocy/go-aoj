package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	solve(os.Stdout, os.Stdin)
}

func solve(dst io.Writer, src io.Reader) {
	var stack stack

	for {
		var s string
		if _, err := fmt.Fscan(src, &s); err != nil {
			break
		}

		switch s {
		case "+":
			last, first := stack.pop(), stack.pop()
			stack.push(first + last)
		case "-":
			last, first := stack.pop(), stack.pop()
			stack.push(first - last)
		case "*":
			last, first := stack.pop(), stack.pop()
			stack.push(first * last)
		default:
			v, err := strconv.Atoi(s)
			if err != nil {
				break
			}
			stack.push(v)
		}
	}

	fmt.Fprintln(dst, stack.pop())
}

type stack []int

func (s *stack) push(v int) {
	*s = append(*s, v)
}

func (s *stack) pop() int {
	if len(*s) == 0 {
		return 0
	}

	lastIdx := len(*s) - 1
	last := (*s)[lastIdx]
	*s = (*s)[:lastIdx]
	return last
}
