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
	list   []T
	size   int // len of list
	length int // len of stack
	top    int // top of stack
	bottom int // bottom of stack
}

func NewStack[T any]() *Stack[T] {
	s := new(Stack[T])

	s.size = defaultSize
	s.list = make([]T, s.size)

	s.length = 0
	s.top = -1
	s.bottom = 0
	return s
}
func NewStackWithSlice[T any](l []T) *Stack[T] {
	s := new(Stack[T])

	s.size = max(len(l), defaultSize)
	s.list = make([]T, s.size)
	copy(s.list, l)

	s.length = s.size
	s.top = s.size - 1
	s.bottom = 0

	return s
}

func (s *Stack[T]) Len() int {
	return s.length
}

func (s *Stack[T]) Cap() int {
	return s.size
}
func (s *Stack[T]) IsEmpty() bool {
	return s.top < s.bottom
}

func (s *Stack[T]) Pop() (r T, ok bool) {
	if s.IsEmpty() {
		return r, false
	}

	s.top--
	s.length--

	return s.list[s.top+1], true

}

func (s *Stack[T]) Push(r T) {
	if s.top+1 == len(s.list) {
		//不使用s.size的原因是 后续可能有 RPush的需求

		s.size *= 2
		newList := make([]T, s.size)
		copy(newList, s.list)
		s.list = newList
	}

	s.list[s.top+1] = r
	s.top++
	s.length++
}
