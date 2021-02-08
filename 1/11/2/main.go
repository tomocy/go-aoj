package main

import (
	"bytes"
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

	b := new(bytes.Buffer)

	var q int
	fmt.Fscan(src, &q)

	for i := 0; i < q; i++ {
		var top, s int
		fmt.Fscan(src, &top, &s)

		dice = dice.rollTo(top, s)
		fmt.Fprintln(b, dice.e)
	}

	fmt.Fprint(dst, b)
}

type dice struct {
	top, bottom int
	n, s, w, e  int
}

func (d dice) rollTo(top, s int) dice {
	rolled := d

	for i := 0; rolled.top != top && i < 3; i++ {
		rolled = rolled.roll(dirN)
	}
	for i := 0; rolled.top != top && i < 3; i++ {
		rolled = rolled.roll(dirW)
	}

	for i := 0; rolled.s != s && i < 3; i++ {
		rolled = rolled.spinLeft()
	}

	return rolled
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

func (d dice) spinLeft() dice {
	return dice{
		top:    d.top,
		bottom: d.bottom,
		n:      d.w,
		s:      d.e,
		w:      d.s,
		e:      d.n,
	}
}

const (
	dirN = 'N'
	dirS = 'S'
	dirW = 'W'
	dirE = 'E'
)
