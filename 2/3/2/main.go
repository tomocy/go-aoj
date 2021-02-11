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
	var n, q int
	fmt.Fscan(src, &n, &q)

	var queue queue
	for i := 0; i < n; i++ {
		var t task
		fmt.Fscan(src, &t.name, &t.time)
		queue.enqueue(t)
	}

	b := new(bytes.Buffer)

	var used int
	for !queue.isEmpty() {
		task := queue.dequeue()

		time, completed := task.proceed(q)
		used += time
		if !completed {
			queue.enqueue(task)
			continue
		}

		fmt.Fprintf(b, "%s %d\n", task.name, used)
	}

	fmt.Fprint(dst, b)
}

type queue []task

func (q *queue) enqueue(t task) {
	*q = append(*q, t)
}

func (q *queue) dequeue() task {
	if q.isEmpty() {
		return task{}
	}

	first := (*q)[0]
	*q = (*q)[1:]
	return first
}

func (q queue) isEmpty() bool {
	return len(q) == 0
}

type task struct {
	name string
	time int
}

func (t *task) proceed(quantum int) (int, bool) {
	used := t.time
	if used > quantum {
		used = quantum
	}

	t.time -= used
	if t.time < 0 {
		t.time = 0
	}

	return used, t.time == 0
}
