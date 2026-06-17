package main

import (
	"fmt"
)

func main() {
	q := NewQueue[string]()

teams := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P"}
	for _, team := range teams {
		q.Enqueue(team)
	}

	for i := 0; i < 15; i++ {
		time1 := q.Dequeue()
		time2 := q.Dequeue()
		var gols1, gols2 int
		if _, err := fmt.Scan(&gols1, &gols2); err != nil {
			return 
	}

	if gols1 > gols2 {
		q.Enqueue(time1)
	} else {
		q.Enqueue(time2)
	}

	}

	fmt.Println(q.Dequeue())
}