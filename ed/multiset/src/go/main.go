package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data     []int
	size     int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet {
	if capacity < 1 {
		capacity = 1
	}
	return &MultiSet{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func (s *MultiSet) String() string {
	return "[" + Join(s.data[:s.size], ", ") + "]"
}

func (s *MultiSet) Reserve(newCapacity int) {
	if newCapacity > s.capacity {
		newData := make([]int, newCapacity)
		for i := 0; i < s.size; i++ {
			newData[i] = s.data[i]
		}
		s.data = newData
		s.capacity = newCapacity
	}
}

func (s *MultiSet) search(value int) (bool, int) {
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

func (s *MultiSet) Insert(value int) {
	_, index := s.search(value)

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

func (s *MultiSet) Erase(value int) error {
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

func (s *MultiSet) Contains(value int) bool {
	found, _ := s.search(value)
	return found
}

func (s *MultiSet) Count(value int) int {
	found, index := s.search(value)
	if !found {
		return 0
	}

	count := 1
	for i := index - 1; i >= 0 && s.data[i] == value; i-- {
		count++
	}
	for i := index + 1; i < s.size && s.data[i] == value; i++ {
		count++
	}

	return count
}

func (s *MultiSet) Unique() int {
	if s.size == 0 {
		return 0
	}

	count := 1
	for i := 1; i < s.size; i++ {
		if s.data[i] != s.data[i-1] {
			count++
		}
	}
	return count
}

func (s *MultiSet) Clear() {
	s.size = 0
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	s := NewMultiSet(0)

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
				s = NewMultiSet(value)
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
		case "count":
			if len(parts) >= 2 {
				value, _ := strconv.Atoi(parts[1])
				fmt.Println(s.Count(value))
			}
		case "unique":
			fmt.Println(s.Unique())
		case "clear":
			s.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}