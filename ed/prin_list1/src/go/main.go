package main

import (
	"fmt"
)

func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	if it.next == l.root {
		return l.root.next
	}
	return it.next
}

func main() {
	var qtd, chosen int

	
	for {
		n, _ := fmt.Scan(&qtd, &chosen)
		if n != 2 {
			break
		}

		l := NewDList[int]()
		for i := 1; i <= qtd; i++ {
			l.PushBack(i)
		}

		sword := l.Front()

		for range chosen - 1 {
			sword = Next(l, sword)
		}

		for range qtd - 1 {
			fmt.Println(ToStr(l, sword))
			l.Erase(Next(l, sword))
			sword = Next(l, sword)
		}

		fmt.Println(ToStr(l, sword))
	}
}