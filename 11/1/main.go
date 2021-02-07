package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	solve(os.Stdout, os.Stdin)
}

func solve(dst io.Writer, src io.Reader) {
	var dice dice
	fmt.Fscan(src, &dice.top, &dice.s, &dice.e, &dice.w, &dice.n, &dice.bottom)

	var dirs string
	fmt.Fscan(src, &dirs)

	for _, dir := range dirs {
		dice = dice.roll(dir)
	}

	fmt.Fprintln(dst, dice.top)
}

type dice struct {
	top, bottom int
	n, s, w, e  int
}

func (d dice) roll(to rune) dice {
	switch to {
	case dirN:
		return dice{
			top:    d.s,
			bottom: d.n,
			n:      d.top,
			s:      d.bottom,
			w:      d.w,
			e:      d.e,
		}
	case dirS:
		return dice{
			top:    d.n,
			bottom: d.s,
			n:      d.bottom,
			s:      d.top,
			w:      d.w,
			e:      d.e,
		}
	case dirW:
		return dice{
			top:    d.e,
			bottom: d.w,
			n:      d.n,
			s:      d.s,
			w:      d.top,
			e:      d.bottom,
		}
	case dirE:
		return dice{
			top:    d.w,
			bottom: d.e,
			n:      d.n,
			s:      d.s,
			w:      d.bottom,
			e:      d.top,
		}
	default:
		panic(fmt.Sprintf("unknown dir: %c", to))
	}
}

const (
	dirN = 'N'
	dirS = 'S'
	dirW = 'W'
	dirE = 'E'
)
