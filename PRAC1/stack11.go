package main

import (
	"errors"
)



type SNode struct {
	next *SNode
	val string
}

type Stack struct {
	head *SNode
}

func (s *Stack) Spush(val string) {
	newNode := &SNode{val: val}
	newNode.next = s.head
	s.head = newNode
}

func (s *Stack) Spop() (string, error) {
	if s.head == nil {
		return "", errors.New("stack is empty")
	}
	val := s.head.val 
	s.head = s.head.next
	return val, errors.New("")
}

func (s *Stack) Peek() (string, error) {//лишнее
	if s.head == nil {
		return "", errors.New("stack is empty")
	}
	return s.head.val, errors.New("")
}

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}

