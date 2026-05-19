package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	root := &Node{}
	root.next = root
	root.prev = root
	return &LList{root: root, size: 0}
}

func (l *LList) insert(it *Node, value int) {
	newNode := &Node{value: value, next: it, prev: it.prev}
	it.prev.next = newNode
	it.prev = newNode
	l.size++
}

func (l *LList) erase(it *Node) {
	if it == l.root {
		return
	}
	it.prev.next = it.next
	it.next.prev = it.prev
	l.size--
}

func (l *LList) PushBack(value int) {
	l.insert(l.root, value)
}

func (l *LList) PushFront(value int) {
	l.insert(l.root.next, value)
}

func (l *LList) PopBack() {
	if l.size > 0 {
		l.erase(l.root.prev)
	}
}

func (l *LList) PopFront() {
	if l.size > 0 {
		l.erase(l.root.next)
	}
}

func (l *LList) Size() int {
	return l.size
}

func (l *LList) Clear() {
	l.root.next = l.root
	l.root.prev = l.root
	l.size = 0
}

func (l *LList) String() string {
	var values []string
	for it := l.root.next; it != l.root; it = it.next {
		values = append(values, strconv.Itoa(it.value))
	}
	return "[" + strings.Join(values, ", ") + "]"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println("$" + line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}