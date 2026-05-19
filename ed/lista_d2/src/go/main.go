package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next  *Node
	prev  *Node
	root  *Node
}

func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}
	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}
	return n.prev
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	root := &Node{}
	root.next = root
	root.prev = root
	root.root = root
	return &LList{root: root, size: 0}
}

func (l *LList) Front() *Node {
	if l.size == 0 {
		return nil
	}
	return l.root.next
}

func (l *LList) Back() *Node {
	if l.size == 0 {
		return nil
	}
	return l.root.prev
}

func (l *LList) Search(value int) *Node {
	for it := l.root.next; it != l.root; it = it.next {
		if it.Value == value {
			return it
		}
	}
	return nil
}

func (l *LList) Insert(it *Node, value int) {
	if it == nil {
		return
	}
	newNode := &Node{Value: value, next: it, prev: it.prev, root: l.root}
	it.prev.next = newNode
	it.prev = newNode
	l.size++
}

func (l *LList) Remove(it *Node) *Node {
	if it == nil || it == l.root {
		return nil
	}

	prox := it.next

	it.prev.next = prox
	prox.prev = it.prev
	l.size--

	if prox == l.root {
		return nil
	}
	return prox
}

func (l *LList) PushBack(value int) {
	l.Insert(l.root, value)
}

func (l *LList) PushFront(value int) {
	l.Insert(l.root.next, value)
}

func (l *LList) PopBack() {
	if l.size > 0 {
		l.Remove(l.root.prev)
	}
}

func (l *LList) PopFront() {
	if l.size > 0 {
		l.Remove(l.root.next)
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
		values = append(values, strconv.Itoa(it.Value))
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
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.Value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}