package main

import (
	"fmt"
	"strings"
)

type DNode[T comparable] struct {
	Value            T
	next, prev, root *DNode[T]
}

type DList[T comparable] struct {
	root *DNode[T]
	size int
}

func NewDList[T comparable]() *DList[T] {
	root := &DNode[T]{}
	root.next = root
	root.prev = root
	root.root = root
	return &DList[T]{root: root, size: 0}
}

func (l *DList[T]) Front() *DNode[T] {
	if l.size == 0 {
		return nil
	}
	return l.root.next
}

func (l *DList[T]) PushBack(value T) {
	l.Insert(l.root, value)
}

func (l *DList[T]) Insert(it *DNode[T], value T) {
	n := &DNode[T]{Value: value, root: l.root}
	n.prev = it.prev
	n.next = it
	it.prev.next = n
	it.prev = n
	l.size++
}

func (l *DList[T]) Erase(it *DNode[T]) {
	if it == l.root || it == nil {
		return
	}
	it.prev.next = it.next
	it.next.prev = it.prev
	l.size--
}

func ToStr(l *DList[int], sword *DNode[int]) string {
	var sb strings.Builder
	sb.WriteString("[")
	for n := l.root.next; n != l.root; n = n.next {
		sb.WriteString(fmt.Sprintf(" %v", n.Value))
		if n == sword {
			sb.WriteString(">")
		}
	}
	sb.WriteString(" ]")
	return sb.String()
}