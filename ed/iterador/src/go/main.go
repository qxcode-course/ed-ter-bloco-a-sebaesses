package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MyList struct {
	data []int
}

type Iterator struct {
	data  []int
	index int
}

type ReverseIterator struct {
	data  []int
	index int
}

func NewMyList(values []int) *MyList {
	return &MyList{data: values}
}

func (l *MyList) Iterator() *Iterator {
	return &Iterator{data: l.data, index: -1}
}

func (i *Iterator) HasNext() bool {
	return i.index < len(i.data)-1
}

func (i *Iterator) Next() int {
	if i.index == len(i.data) {
		panic(fmt.Errorf("No more elements"))
	}
	i.index += 1
	return i.data[i.index]
}

func (l *MyList) ReverseIterator() *ReverseIterator {
	return &ReverseIterator{data: l.data, index: len(l.data)}
}

func (i *ReverseIterator) HasNext() bool {
	return i.index > 0
}

func (i *ReverseIterator) Next() int {
	if i.index == 0 {
		panic(fmt.Errorf("No more elements"))
	}
	i.index--
	return i.data[i.index]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	mylist := NewMyList([]int{})

	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			slice := make([]int, 0, len(args)-1)
			for _, value := range args[1:] {
				val, _ := strconv.Atoi(value)
				slice = append(slice, val)
			}
			mylist = NewMyList(slice)
		case "show":
			fmt.Print("[ ")
			for it := mylist.Iterator(); it.HasNext(); {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		case "reverse":
			fmt.Print("[ ")
			for it := mylist.ReverseIterator(); it.HasNext(); {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		case "cyclic":
		}
	}
}