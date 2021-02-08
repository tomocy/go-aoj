package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	solve(os.Stdout, os.Stdin)
}

func solve(dst io.Writer, src io.Reader) {
	input, err := ioutil.ReadAll(src)
	if err != nil {
		return
	}
	src = bytes.NewReader(input)

	var n int
	fmt.Fscan(src, &n)

	rates := make([]int, n)
	for i := range rates {
		fmt.Fscan(src, &rates[i])
	}

	minRate, maxProfit := rates[0], rates[1]-rates[0]
	for i := 1; i < len(rates); i++ {
		if profit := rates[i] - minRate; profit > maxProfit {
			maxProfit = profit
		}
		if rates[i] < minRate {
			minRate = rates[i]
		}
	}

	fmt.Fprintln(dst, maxProfit)
}
