package stack

// Copyright (c) 2024, Lee.
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.
const (
	defaultSize = 100
)

type Stack[T any] struct {
	list []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		list: make([]T, 0, defaultSize),
	}
}

func NewStackWithSlice[T any](l []T) *Stack[T] {
	s := new(Stack[T])

	s.list = make([]T, defaultSize)
	copy(s.list, l)

	return s
}

// Push adds an item to the top of the stack.
func (s *Stack[T]) Push(item T) {
	s.list = append(s.list, item)
}

// Pop removes the item on the top of the stack and returns it.
// If the stack is empty, Pop returns the zero value of the type T and an error.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.list) == 0 {
		var zero T // Create a zero value of type T
		return zero, false
	}
	index := len(s.list) - 1
	item := s.list[index]
	s.list = s.list[:index]
	return item, true
}

// Top returns the item on the top of the stack without removing it.
// If the stack is empty, Top returns the zero value of the type T and an error.
func (s *Stack[T]) Top() (T, bool) {
	if len(s.list) == 0 {
		var zero T // Create a zero value of type T
		return zero, false
	}
	return s.list[len(s.list)-1], true
}

// IsEmpty returns true if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.list) == 0
}

// Size returns the number of items in the stack.
func (s *Stack[T]) Size() int {
	return len(s.list)
}

// Clear removes all items from the stack.
func (s *Stack[T]) Clear() {
	s.list = nil // Clearing the slice
}
