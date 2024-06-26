package Stack

// Stack represents a stack data structure
type Stack struct {
	elements []any
}

// InitStack initializes the stack
func InitStack() *Stack {
	return &Stack{}
}

// Push adds an element to the top of the stack
func (s *Stack) Push(value any) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top element from the stack
func (s *Stack) Pop() any {
	if s.IsEmpty() {
		return nil
	}

	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return top
}

// Peek returns the top element of the stack without removing it
func (s *Stack) Peek() any {
	if s.IsEmpty() {
		return nil
	}

	return s.elements[len(s.elements)-1]
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

// Display displays the stack
func (s *Stack) Display() {
	for i := len(s.elements) - 1; i >= 0; i-- {
		print(s.elements[i], " ")
	}
	println()
}

// Length returns the number of elements in the stack
func (s *Stack) Length() int {
	return len(s.elements)
}

// Clear removes all elements from the stack
func (s *Stack) Clear() {
	s.elements = []any{}
}
