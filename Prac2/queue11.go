package main

import "errors"

// Queue represents a queue data structure.
type Queue struct {
	items []string
}

// Qadd adds an element to the end of the queue.
func (q *Queue) Qadd(item string) {
	q.items = append(q.items, item)
}

func (q *Queue) Qdell() (string, error) {
	if len(q.items) == 0 {
		return "", errors.New("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, errors.New("")
}

// IsEmpty returns true if the queue is empty, and false otherwise.
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Qsize() int {
	return len(q.items)
}
