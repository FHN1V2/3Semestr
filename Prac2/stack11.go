package main

import "sync"

type SNode struct {
	next *SNode
	val  string
}

type Stack struct {
	mu   sync.Mutex
	head *SNode
}

func (s *Stack) Spush(val string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	newNode := &SNode{val: val}
	newNode.next = s.head
	s.head = newNode
}

func (s *Stack) Spop() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.head == nil {
		return "error"
	}
	val := s.head.val
	s.head = s.head.next
	return val
}

func (s *Stack) Peek() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.head == nil {
		return ""
	}
	return s.head.val
}

func (s *Stack) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.head == nil
}
