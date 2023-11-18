package main

import "fmt"

func main() {
	fmt.Println("")
	return
}

type Stack struct {
	items []int
	size  int
}

func (s *Stack) addToStack(item int) error {
	if s.size == s.len() {
		return fmt.Errorf("Стек переполнен")
	}
	s.items = append(s.items, item)
	return nil
}

func (s *Stack) getFromStack() (int, error) {
	if s.isEmpty() {
		return 0, fmt.Errorf("Стек пуст")
	}
	item := s.items[s.len()-1]
	if s.len() == 1 {
		s.items = nil
		return item, nil
	}
	s.items = s.items[:s.len()-1]
	return item, nil
}

func (s *Stack) isEmpty() bool {
	return s.len() == 0
}

func (s *Stack) len() int {
	return len(s.items)
}
