package stack

import "container/list"

// Stack is a stack structure implemented on top of a linked list.
type Stack struct {
	storage *list.List
}

// New returns an initialized stack.
func New() *Stack {
	return new(Stack).Init()
}

// Init initializes our stack through its internal storage (linked list).
func (s *Stack) Init() *Stack {
	s.storage = list.New()
	return s
}

// Len returns the length of the stack
func (s *Stack) Len() int {
	return s.storage.Len()
}

// Push adds an item with the given value to the end of the stack.
func (s *Stack) Push(v interface{}) {
	s.storage.PushBack(v)
}

// Pop returns and removes the value of the last item in the stack.
func (s *Stack) Pop() interface{} {
	el := s.storage.Back()
	return s.storage.Remove(el)
}

// Top returns the value of the last item in the stack without removing it.
func (s *Stack) Top() interface{} {
	el := s.storage.Back()
	return el.Value
}
