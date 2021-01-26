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
	var row, col int
	fmt.Fscanln(src, &row, &col)

	table := make([][]int, row, row+1)
	for i := range table {
		table[i] = make([]int, col, col+1)
	}

	for i, row := range table {
		for j := range row {
			var v int
			fmt.Fscan(src, &v)
			table[i][j] = v
		}

		fmt.Fscan(src)
	}

	colSums := make([]int, col)
	for i := range colSums {
		for j := range table {
			colSums[i] += table[j][i]
		}
	}
	table = append(table, colSums)

	for i, row := range table {
		var sum int
		for _, v := range row {
			sum += v
		}

		table[i] = append(table[i], sum)
	}

	b := new(strings.Builder)

	for _, row := range table {
		for j, v := range row {
			if j != 0 {
				fmt.Fprint(b, " ")
			}

			fmt.Fprint(b, v)
		}

		fmt.Fprintln(b)
	}

	fmt.Fprint(dst, b)
}
