package main

import "sync"


// Queue represents a queue data structure.
type Queue struct {
	mu    sync.Mutex
	items []string
}

// Qadd adds an element to the end of the queue.
func (q *Queue) Qpush(item string) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
}

func (q *Queue) Qpop() string {
	if len(q.items) == 0 {
		return "Error"
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// IsEmpty returns true if the queue is empty, and false otherwise.
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Qsize() int {
	return len(q.items)
}
