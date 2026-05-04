package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(capacity int) *Set {
	if capacity < 1 {
		capacity = 1
	}
	return &Set{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (s *Set) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < s.size; i++ {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(strconv.Itoa(s.data[i]))
	}
	sb.WriteString("]")
	return sb.String()
}

func (s *Set) Reserve(newCapacity int) {
	if newCapacity > s.capacity {
		newData := make([]int, newCapacity)
		for i := 0; i < s.size; i++ {
			newData[i] = s.data[i]
		}
		s.data = newData
		s.capacity = newCapacity
	}
}


func (s *Set) search(value int) (bool, int) {
	low := 0
	high := s.size

	for low < high {
		mid := low + (high-low)/2

		if s.data[mid] == value {
			return true, mid
		} else if s.data[mid] < value {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return false, low
}

func (s *Set) Insert(value int) {
	found, index := s.search(value)

	
	if found {
		return
	}

	
	if s.size == s.capacity {
		newCap := s.capacity * 2
		if newCap == 0 {
			newCap = 1
		}
		s.Reserve(newCap)
	}

	
	for i := s.size; i > index; i-- {
		s.data[i] = s.data[i-1]
	}
	
	s.data[index] = value
	s.size++
}

func (s *Set) Erase(value int) error {
	found, index := s.search(value)

	
	if !found {
		return fmt.Errorf("value not found")
	}

	
	for i := index; i < s.size-1; i++ {
		s.data[i] = s.data[i+1]
	}
	s.size--
	return nil
}

func (s *Set) Contains(value int) bool {
	found, _ := s.search(value)
	return found
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	s := NewSet(0)

	for {
		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
		fmt.Println("$" + line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			if len(parts) >= 2 {
				value, _ := strconv.Atoi(parts[1])
				s = NewSet(value)
			}
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				s.Insert(value)
			}
		case "show":
			fmt.Println(s)
		case "erase":
			if len(parts) >= 2 {
				value, _ := strconv.Atoi(parts[1])
				err := s.Erase(value)
				if err != nil {
					fmt.Println(err)
				}
			}
		case "contains":
			if len(parts) >= 2 {
				value, _ := strconv.Atoi(parts[1])
				if s.Contains(value) {
					fmt.Println("true")
				} else {
					fmt.Println("false")
				}
			}
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}