package main

import (
	"fmt"
	"sync"
	"strings"

)

type Set struct {
	mu sync.Mutex
	data map[string]bool
}

func NewSet() *Set {
	return &Set{data: make(map[string]bool)}
}

func (s *Set) SetAdd(item string) {
	s.mu.Lock()
	s.data[item] = true
	defer s.mu.Unlock()
	
}

func (s *Set) SetRemove(item string) {
	if a := s.data[item]; !a {
		fmt.Println("Not found")
	}

	newData := make(map[string]bool)
	for key := range s.data {
		if key != item {
			newData[key] = false
		}
	}
	s.data = newData
}

func (s *Set) SetContains(item string) bool {
	return s.data[item]
}

func (s *Set) SetSize() int {
	return len(s.data)
}

func (s *Set) SetPrint() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	var items []string
	for item := range s.data {
		items = append(items, item)
	}

	return "Set content: " + strings.Join(items, ", ")
}