package monotoneStack

import "golang.org/x/exp/constraints"

// Copyright (c) 2024, Lee.
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.
const (
	defaultSize = 100
)

const (
	monotoneTypeCustomize MonotoneType = iota

	MonotoneTypeIncrease
	MonotoneTypeDecrease
)

type MonotoneType int

type MonotoneStack[T any] struct {
	list []T // stack

	// t is the type of monotone stack
	//
	// monotoneTypeCustomize is used in NewMonotoneStackWithCheck
	// MonotoneTypeIncrease Monotonically increasing stack, default monotone type
	// MonotoneTypeDecrease Monotonically decreasing stack
	t MonotoneType

	// strict controls whether the stack is strictly monotonic
	//
	// when t is monotoneTypeCustomize, strict is invalid
	//
	// default strict = false
	strict bool

	// compare is Custom comparison function
	//
	// when it returns true, elem is pushed into the stack;
	// when it returns false, top is popped out of the stack.
	// This function is used to determine the monotonicity when data is pushed into the stack
	compare compareFunc[T]
}

func NewMonotoneStack[T constraints.Ordered](t MonotoneType, isStrict bool) *MonotoneStack[T] {
	if t != MonotoneTypeIncrease && t != MonotoneTypeDecrease {
		t = MonotoneTypeIncrease
	}
	return &MonotoneStack[T]{
		list:    make([]T, 0, defaultSize),
		t:       t,
		strict:  isStrict,
		compare: compareGet[T](t, isStrict),
	}
}

func NewMonotoneStackWithCompare[T any](f func(top, elem T) bool) *MonotoneStack[T] {
	return &MonotoneStack[T]{
		list:    make([]T, 0, defaultSize),
		t:       monotoneTypeCustomize,
		compare: f,
	}
}

// Push adds an elem to the top of the stack.
//
// return the popList those popped from the stack due to elem being pushed onto stack
//
// popList keeps the order of popping
func (s *MonotoneStack[T]) Push(elem T) []T {
	popList := make([]T, 0)
	for !s.check(elem) {
		if top, ok := s.pop(); ok {
			popList = append(popList, top)
		}
	}
	s.list = append(s.list, elem)
	return popList
}

// TryPush attempts to push elem onto the monotone stack.
//
// return whether elem is pushed onto the stack successfully
func (s *MonotoneStack[T]) TryPush(elem T) bool {
	if s.check(elem) {
		s.list = append(s.list, elem)
		return true
	}
	return false
}

// Check whether the elem can be pushed onto the stack
func (s *MonotoneStack[T]) Check(elem T) bool {
	return s.check(elem)
}

// Pop removes the item on the top of the stack and returns it.
// If the stack is empty, Pop returns the zero value of the type T and an error.
func (s *MonotoneStack[T]) Pop() (T, bool) {
	return s.pop()
}

// Top returns the item on the top of the stack without removing it.
// If the stack is empty, Top returns the zero value of the type T and an error.
func (s *MonotoneStack[T]) Top() (T, bool) {
	return s.top()
}

// IsEmpty returns true if the stack is empty.
func (s *MonotoneStack[T]) IsEmpty() bool {
	return s.isEmpty()
}

// Size returns the number of items in the stack.
func (s *MonotoneStack[T]) Size() int {
	return len(s.list)
}

// Clear removes all items from the stack.
func (s *MonotoneStack[T]) Clear() {
	s.list = nil // Clearing the slice
}

// List return the copy of stack.
func (s *MonotoneStack[T]) List() []T {
	stackCopy := make([]T, len(s.list))
	copy(stackCopy, s.list)
	return stackCopy
}

// top returns the item on the top of the stack without removing it.
// If the stack is empty, Top returns the zero value of the type T and an error.
func (s *MonotoneStack[T]) top() (T, bool) {
	if len(s.list) == 0 {
		var zero T // Create a zero value of type T
		return zero, false
	}
	return s.list[len(s.list)-1], true
}

// isEmpty returns true if the stack is empty.
func (s *MonotoneStack[T]) isEmpty() bool {
	return len(s.list) == 0
}

// pop removes the item on the top of the stack and returns it.
// If the stack is empty, Pop returns the zero value of the type T and an error.
func (s *MonotoneStack[T]) pop() (T, bool) {
	if len(s.list) == 0 {
		var zero T // Create a zero value of type T
		return zero, false
	}
	index := len(s.list) - 1
	item := s.list[index]
	s.list = s.list[:index]
	return item, true
}

// check whether the elem can be pushed onto the stack
func (s *MonotoneStack[T]) check(elem T) bool {
	if top, ok := s.top(); ok {
		return s.compare(top, elem)
	}
	return true
}
