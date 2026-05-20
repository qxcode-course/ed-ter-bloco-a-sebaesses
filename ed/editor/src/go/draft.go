package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node[T comparable] struct {
	Value T
	next  *Node[T]
	prev  *Node[T]
	root  *Node[T]
}

type LList[T comparable] struct {
	root *Node[T]
	size int
}

func NewLList[T comparable]() *LList[T] {
	root := &Node[T]{}
	root.next = root
	root.prev = root
	root.root = root
	return &LList[T]{root: root, size: 0}
}

func (n *Node[T]) Next() *Node[T] {
	if n.next == n.root {
		return n.root.next
	}
	return n.next
}

func (l *LList[T]) Size() int {
	return l.size
}

func (l *LList[T]) Clear() {
	l.root.next = l.root
	l.root.prev = l.root
	l.size = 0
}

func (l *LList[T]) PushBack(value T) {
	l.insertBefore(l.root, value)
}

func (l *LList[T]) insertBefore(mark *Node[T], value T) {
	n := &Node[T]{Value: value, root: l.root}
	n.prev = mark.prev
	n.next = mark
	mark.prev.next = n
	mark.prev = n
	l.size++
}

func (l *LList[T]) String() string {
	values := []string{}
	for n := l.root.next; n != l.root; n = n.next {
		values = append(values, fmt.Sprint(n.Value))
	}
	return "[" + strings.Join(values, ", ") + "]"
}

func (l *LList[T]) Search(value T) *Node[T] {
	for n := l.root.next; n != l.root; n = n.next {
		if n.Value == value {
			return n
		}
	}
	return nil
}

func (l *LList[T]) removeNode(mark *Node[T]) {
	if mark == l.root {
		return
	}
	mark.prev.next = mark.next
	mark.next.prev = mark.prev
	l.size--
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		editor := NewLList[int]()
		cursor := editor.root

		for _, c := range line {
			switch c {
			case 'R':
				editor.insertBefore(cursor, int('\n'))
			case 'B':
				if cursor.prev != editor.root {
					editor.removeNode(cursor.prev)
				}
			case 'D':
				if cursor != editor.root {
					toDelete := cursor
					cursor = cursor.next
					editor.removeNode(toDelete)
				}
			case '<':
				if cursor.prev != editor.root {
					cursor = cursor.prev
				}
			case '>':
				if cursor != editor.root {
					cursor = cursor.next
				}
			default:
				editor.insertBefore(cursor, int(c))
			}
		}

		for n := editor.root.next; n != editor.root; n = n.next {
			if n == cursor {
				fmt.Print("|")
			}
			fmt.Print(string(rune(n.Value)))
		}

		if cursor == editor.root {
			fmt.Print("|")
		}
		fmt.Println()
	}
}