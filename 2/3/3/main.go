package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	solve(os.Stdout, os.Stdin)
}

func solve(dst io.Writer, src io.Reader) {
	src = bufio.NewReader(src)

	scanner := bufio.NewScanner(src)

	scanner.Scan()

	var list list
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		cmd := s[0]
		switch cmd {
		case "insert":
			v, _ := strconv.Atoi(s[1])
			list.insert(v)
		case "delete":
			v, _ := strconv.Atoi(s[1])
			list.delete(v)
		case "deleteFirst":
			list.deleteFirst()
		case "deleteLast":
			list.deleteLast()
		}
	}

	b := new(bytes.Buffer)

	list.doEach(func(i int, item *listItem) bool {
		if i != 0 {
			fmt.Fprint(b, " ")
		}
		fmt.Fprint(b, item.v)
		return false
	})
	fmt.Fprintln(b)

	fmt.Fprint(dst, b)
}

type list struct {
	dummy *listItem
}

func (l *list) insert(v int) {
	l.init()

	l.insertFirst(v)
}

func (l *list) insertFirst(v int) {
	l.init()

	item := &listItem{
		v:    v,
		prev: l.dummy,
		next: l.dummy.next,
	}
	l.dummy.next.prev = item
	l.dummy.next = item
}

func (l *list) delete(v int) {
	l.init()
	l.doEach(func(_ int, item *listItem) bool {
		if item.v == v {
			l.deleteItem(item)
			return true
		}
		return false
	})
}

func (l *list) deleteFirst() {
	l.init()
	l.deleteItem(l.first())
}

func (l *list) deleteLast() {
	l.init()
	l.deleteItem(l.last())
}

func (l *list) deleteItem(item *listItem) {
	l.init()

	if item == l.dummy {
		return
	}

	item.prev.next = item.next
	item.next.prev = item.prev
}

func (l *list) doEach(do func(int, *listItem) bool) {
	l.init()

	for i, item := 0, l.first(); item != l.end(); i, item = i+1, item.next {
		if do(i, item) {
			break
		}
	}
}

func (l *list) first() *listItem {
	l.init()
	return l.dummy.next
}

func (l *list) last() *listItem {
	l.init()
	return l.dummy.prev
}

func (l *list) end() *listItem {
	l.init()
	return l.dummy
}

func (l *list) init() {
	if l.dummy == nil {
		l.dummy = new(listItem)
		l.dummy.prev, l.dummy.next = l.dummy, l.dummy
	}
}

type listItem struct {
	v    int
	prev *listItem
	next *listItem
}
