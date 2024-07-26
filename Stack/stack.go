package Stack

import "errors"

// Stack represents a stack data structure
type Stack[T any] struct {
	elements []T
}

// InitStack initializes the stack
func InitStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top element from the stack
// Throws error if stack is empty
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var emptyValue T
		return emptyValue, errors.New("Stack is Empty. No elements to POP")
	}

	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return top, nil
}

// Peek returns the top element of the stack without removing it
// Throws error if the stack is empty
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var emptyValue T
		return emptyValue, errors.New("Stack is Empty")
	}

	return s.elements[len(s.elements)-1], nil
}

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Display displays the stack
func (s *Stack[T]) Display() {
	for i := len(s.elements) - 1; i >= 0; i-- {
		print(s.elements[i], " ")
	}
	println()
}

// Length returns the number of elements in the stack
func (s *Stack[T]) Length() int {
	return len(s.elements)
}

// Clear removes all elements from the stack
func (s *Stack[T]) Clear() {
	s.elements = s.elements[:0]
}
