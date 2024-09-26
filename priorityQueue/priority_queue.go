package priorityQueue

import "golang.org/x/exp/constraints"

const (
	defaultSize = 100
)

const (
	priorityTypeCustomize PriorityType = iota

	PriorityTypeGreater
	PriorityTypeLower
)

type PriorityType int

type PriorityQueue[T any] struct {
	list []T // queue

	// t is the type of priority queue
	//
	// priorityTypeCustomize is used in NewMonotoneStackWithCheck
	// PriorityTypeGreater Greater value first, default priority type
	// PriorityTypeLower Lower value first
	t PriorityType

	// compare is Custom comparison function
	//
	// compare src and des :
	//
	// when priority of src higher than priority of des returns true;
	//
	// This function is used to determine priority of elems
	compare compareFunc[T]
}

// NewPriorityQueue returns a new PriorityQueue.
//
// t is the type of priority queue : greater or lower
func NewPriorityQueue[T constraints.Ordered](t PriorityType) *PriorityQueue[T] {
	if t != PriorityTypeGreater && t != PriorityTypeLower {
		t = PriorityTypeGreater
	}
	return &PriorityQueue[T]{
		list:    make([]T, 0, defaultSize),
		t:       t,
		compare: compareGet[T](t),
	}
}

// NewPriorityQueueWithCompare returns a new PriorityQueue.
//
// f is Custom comparison function
//
// compareFunc[T] func(src, des T) bool
//
// when src has a higher priority than des, returns true
func NewPriorityQueueWithCompare[T any](f compareFunc[T]) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		list:    make([]T, 0, defaultSize),
		t:       priorityTypeCustomize,
		compare: f,
	}
}

// IsEmpty returns true if the stack is empty.
func (q *PriorityQueue[T]) IsEmpty() bool {
	return q.isEmpty()
}

// Size returns the number of items in the stack.
func (q *PriorityQueue[T]) Size() int {
	return q.size()
}

// isEmpty returns true if the stack is empty.
func (q *PriorityQueue[T]) isEmpty() bool {
	return len(q.list) == 0
}

// size returns the number of items in the stack.
func (q *PriorityQueue[T]) size() int {
	return len(q.list)
}

// heapIfy(堆化/sink) keep slice as a heap form index 'root'
func (q *PriorityQueue[T]) heapIfy(root int) {
	num := q.size()
	l, r, largest := 2*root+1, 2*root+2, root
	if l < num && q.compare(q.list[l], q.list[largest]) {
		largest = l
	}

	if r < num && q.compare(q.list[r], q.list[largest]) {
		largest = r
	}
	//若root不是最大值，则交换root与最大的子节点，子节点继续堆化
	if largest != root {
		q.list[root], q.list[largest] = q.list[largest], q.list[root]
		q.heapIfy(largest)
	}

}

// build a heap with slice
func (q *PriorityQueue[T]) build() {
	num := q.size()
	for i := num / 2; i >= 0; i-- {
		q.heapIfy(i)
	}
}

// float keep slice as a heap form index 'last'
func (q *PriorityQueue[T]) float(last int) {
	for last > 0 {
		parent := (last - 1) / 2
		if q.compare(q.list[last], q.list[parent]) {
			q.list[last], q.list[parent] = q.list[parent], q.list[last]
			last = parent
		} else {
			break
		}
	}
}

// Push adds an element to the priority queue.
func (q *PriorityQueue[T]) Push(e T) {
	q.list = append(q.list, e)
	q.float(q.size() - 1)
}

// Pop removes the element with the highest priority from the priority queue and returns it.
// If the priority queue is empty, Pop returns the zero value of the type T and false.
func (q *PriorityQueue[T]) Pop() (T, bool) {
	if q.isEmpty() {
		var zero T
		return zero, false
	}

	num := q.size() - 1
	e := q.list[0]
	q.list[0], q.list[num] = q.list[num], q.list[0]
	q.list = q.list[:num]
	q.heapIfy(0)
	return e, true
}

// Top returns the element with the highest priority from the priority queue without removing it.
// If the priority queue is empty, Top returns the zero value of the type T and false.
func (q *PriorityQueue[T]) Top() (T, bool) {
	if q.isEmpty() {
		var zero T
		return zero, false
	}
	return q.list[0], true
}
