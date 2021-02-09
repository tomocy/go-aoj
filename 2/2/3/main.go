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
	var n int
	fmt.Fscan(src, &n)

	cards := make([]card, n)
	for i := range cards {
		var s string
		fmt.Fscan(src, &s)
		fmt.Sscanf(s, "%c%d", &cards[i].mark, &cards[i].val)
	}

	bubbleSorted, selectionSorted := make([]card, n), make([]card, n)
	copy(bubbleSorted, cards)
	copy(selectionSorted, cards)

	bubbleSort(bubbleSorted)
	selectionSort(selectionSorted)

	stable := true
	for i := range bubbleSorted {
		if bubbleSorted[i] != selectionSorted[i] {
			stable = false
		}
	}

	b := new(bytes.Buffer)

	printCards(b, bubbleSorted...)
	fmt.Fprintln(b, "Stable")
	printCards(b, selectionSorted...)
	if stable {
		fmt.Fprintln(b, "Stable")
	} else {
		fmt.Fprintln(b, "Not stable")
	}

	fmt.Fprint(dst, b)
}

func bubbleSort(cards []card) {
	var sorted bool
	for !sorted {
		sorted = true
		for i := len(cards) - 1; i >= 1; i-- {
			if cards[i].val < cards[i-1].val {
				cards[i-1], cards[i] = cards[i], cards[i-1]
				sorted = false
			}
		}
	}
}

func selectionSort(cards []card) {
	for i := range cards {
		minIdx := i
		for j := i; j < len(cards); j++ {
			if cards[j].val < cards[minIdx].val {
				minIdx = j
			}
		}

		cards[i], cards[minIdx] = cards[minIdx], cards[i]
	}
}

func printCards(dst io.Writer, cards ...card) {
	for i, card := range cards {
		if i != 0 {
			fmt.Fprint(dst, " ")
		}
		fmt.Fprint(dst, card)
	}
	fmt.Fprintln(dst)
}

type card struct {
	mark rune
	val  int
}

func (c card) String() string {
	return fmt.Sprintf("%c%d", c.mark, c.val)
}
