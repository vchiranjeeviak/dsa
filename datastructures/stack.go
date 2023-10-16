package datastructures

import (
	"errors"
	"fmt"
	"math"
)

type stack struct {
	items []int
}

func (s *stack) getStackLength() int {
	// Returns the current length of the stack
	return len(s.items)
}

func (s *stack) pushItem(item int) int {
	// Pushes the supplied number into the stack and returns the currecnt length of the stack
	// O(1) time complexity
	s.items = append(s.items, item)
	return s.getStackLength()
}

func (s *stack) getTopItem() (int, error) {
	// Returns the top most item of the stack
	// O(1) time complexity
	if s.getStackLength() == 0 {
		return int(math.Inf(-1)), errors.New("Stack is currently empty, no top item.")
	}
	return s.items[s.getStackLength()-1], nil
}

func (s *stack) popItem() (int, error) {
	// Pops the top item of the stack and returns the value
	// O(1) time complexity
	poppingItem, err := s.getTopItem()
	if err != nil {
		return int(math.Inf(-1)), errors.New("Stack is currently empty, no items to pop.")
	}
	s.items = s.items[:s.getStackLength()-1]
	return poppingItem, nil
}

func (s *stack) searchItem(item int) int {
	// Returns the position of the item from top if found, else returns -1
    // O(n) time complexity where n is the number of items in the stack
	for i := s.getStackLength() - 1; i >= 0; i-- {
		if s.items[i] == item {
			return s.getStackLength() - i - 1
		}
	}
	return -1
}

func (s *stack) printStack() {
	for i := s.getStackLength() - 1; i >= 0; i-- {
		if i == s.getStackLength()-1 {
			fmt.Printf("%d -> top\n", s.items[i])
		} else {
			fmt.Println(s.items[i])
		}
	}
}
