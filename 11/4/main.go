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
	var n int
	fmt.Fscan(src, &n)

	dices := make([]dice, n)
	for i := range dices {
		fmt.Fscan(src, &dices[i].top, &dices[i].s, &dices[i].e, &dices[i].w, &dices[i].n, &dices[i].bottom)
	}

	if areAllDifferent(dices...) {
		fmt.Fprintln(dst, "Yes")
	} else {
		fmt.Fprintln(dst, "No")
	}
}

func areAllDifferent(ds ...dice) bool {
	if len(ds) == 1 {
		return true
	}

	target, rest := ds[0], ds[1:]
	for _, d := range rest {
		if target.equal(d) {
			return false
		}
	}

	return areAllDifferent(rest...)
}

type dice struct {
	top, bottom int
	n, s, w, e  int
}

func (d dice) equal(other dice) bool {
	unique := d.uniqueFace()
	d, _ = d.rollToTop(unique)
	other, ok := other.rollTo(d.top, d.s)
	if !ok {
		return false
	}

	return d == other
}

func (d dice) uniqueFace() int {
	duplicated := make(map[int]int)
	duplicated[d.top]++
	duplicated[d.bottom]++
	duplicated[d.n]++
	duplicated[d.s]++
	duplicated[d.w]++
	duplicated[d.e]++

	for face, times := range duplicated {
		if times == 1 {
			return face
		}
	}

	return d.top
}

func (d dice) rollTo(top, s int) (dice, bool) {
	rolled, ok := d.rollToTop(top)
	if !ok {
		return d, false
	}

	for i := 0; rolled.s != s && i < 3; i++ {
		rolled = rolled.spinLeft()
	}

	if rolled.top == top && rolled.s == s {
		return rolled, true
	}

	return d, false
}

func (d dice) rollToTop(top int) (dice, bool) {
	rolled := d

	for i := 0; rolled.top != top && i < 3; i++ {
		rolled = rolled.roll(dirN)
	}
	for i := 0; rolled.top != top && i < 3; i++ {
		rolled = rolled.roll(dirW)
	}

	if rolled.top == top {
		return rolled, true
	}

	return d, false
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
