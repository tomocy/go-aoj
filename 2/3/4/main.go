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

	scanner := bufio.NewScanner(src)

	scanner.Scan()
	s := scanner.Text()
	var (
		downIs stack
		areas  areas
	)
	for i, r := range s {
		switch r {
		case '\\':
			downIs.push(i)
		case '/':
			downI, ok := downIs.pop()
			if !ok {
				continue
			}
			areas.add(area{
				downI: downI,
				area:  i - downI,
			})
		}
	}

	var sum int
	for _, area := range areas {
		sum += area.area
	}

	b := new(bytes.Buffer)

	fmt.Fprintln(b, sum)

	fmt.Fprint(b, len(areas))
	for _, area := range areas {
		fmt.Fprint(b, " ")
		fmt.Fprint(b, area.area)
	}
	fmt.Fprintln(b)

	fmt.Fprint(dst, b)
}

type areas []area

func (as *areas) add(a area) {
	*as = append(*as, a)
	as.combine()
}

func (as *areas) combine() {
	for i := len(*as) - 1; i >= 1; i-- {
		if prevI := i - 1; (*as)[i].downI < (*as)[prevI].downI {
			(*as)[prevI].downI = (*as)[i].downI
			(*as)[prevI].area += (*as)[i].area
			*as = append((*as)[:i], (*as)[i+1:]...)
		}
	}
}

type area struct {
	downI int
	area  int
}

type stack []int

func (s *stack) push(i int) {
	*s = append(*s, i)
}

func (s *stack) pop() (int, bool) {
	if s.isEmpty() {
		return 0, false
	}

	lastI := len(*s) - 1
	last := (*s)[lastI]
	*s = (*s)[:lastI]
	return last, true
}

func (s stack) isEmpty() bool {
	return len(s) == 0
}
