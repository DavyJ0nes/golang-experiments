package stack

import "errors"

type Stack struct {
	slice []int
}

var (
	ErrEmptyStack = errors.New("Stack is Empty")
)

// Push Adds an Item to the Top of the Stack
func (s *Stack) Push(i int) {
	s.slice = append(s.slice, i)
}

// Pop Removes the Top Item From the Stack
func (s *Stack) Pop() (int, error) {
	if len(s.slice) == 0 {
		return 0, ErrEmptyStack
	}

	// Getting Last Element in Slice and Removing it
	val := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]

	return val, nil
}

// Peek Looks at the Top Item in the Stack but Doesn't Return it
func (s *Stack) Peek() (int, error) {
	if len(s.slice) == 0 {
		return 0, ErrEmptyStack
	}

	// Getting Last Element in Slice
	val := s.slice[len(s.slice)-1]

	return val, nil
}
