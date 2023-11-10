package main

import "errors"

type Qnode struct {
	next *Qnode
	val  string
}

// Queue struct
type Queue struct {
	head *Qnode
	tail *Qnode
}

// Qadd adds an element to the end of the queue.
func (q *Queue) Qadd(item string) {
	newNode := &Qnode{
		val:  item,
		next: nil,
	}

	if q.tail == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
}

// Qpop removes and returns the element from the beginning of the queue.
func (q *Queue) Qpop() (string, error) {
	if q.IsEmpty() {
		return "", errors.New("Queue is empty")
	}

	item := q.head.val
	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}
	return item, errors.New("")
}

// Qdell removes and returns the element from tail
func (q *Queue) Qdell() (string, error) {
	if q.IsEmpty() {
		return "", errors.New("Queue is empty")
	}

	item := q.head.val
	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}
	return item, errors.New("")
}

// IsEmpty returns true if the queue is empty,else return false
func (q *Queue) IsEmpty() bool {
	return q.head == nil
}