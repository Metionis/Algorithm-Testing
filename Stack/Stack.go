package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []interface{}
}

// Push adds an element to the top of the stack
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element from the stack
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, nil
}

// Peek returns the top element from the stack without removing it
func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("Stack:", stack.items) // Output: [1 2 3]
	fmt.Println("Size of stack:", stack.Size()) // Output: 3

	top, _ := stack.Peek()
	fmt.Println("Top of stack:", top) // Output: 3

	popped, _ := stack.Pop()
	fmt.Println("Popped element:", popped) // Output: 3

	fmt.Println("Stack after pop:", stack.items) // Output: [1 2]
}
