package queue

const (
	defaultSize = 100
)

// Queue represents a generic FIFO queue.
type Queue[T any] struct {
	list []T
}

// NewQueue creates a new queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// NewQueueWithSlice creates a new queue with the given slice.
func NewQueueWithSlice[T any](l []T) *Queue[T] {
	s := new(Queue[T])

	s.list = make([]T, defaultSize)
	copy(s.list, l)

	return s
}

// Enqueue adds an item to the end of the queue.
func (q *Queue[T]) Enqueue(item T) {
	q.list = append(q.list, item)
}

// Dequeue removes the item from the start of the queue and returns it.
// If the queue is empty, Dequeue returns the zero value of the type T and false.
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.list) == 0 {
		var zero T
		return zero, false
	}
	item := q.list[0]
	q.list = q.list[1:]
	return item, true
}

// Front returns the item at the start of the queue without removing it.
// If the queue is empty, Front returns the zero value of the type T and false.
func (q *Queue[T]) Front() (T, bool) {
	if len(q.list) == 0 {
		var zero T
		return zero, false
	}
	return q.list[0], true
}

// IsEmpty checks if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.list) == 0
}

// Size returns the number of items in the queue.
func (q *Queue[T]) Size() int {
	return len(q.list)
}

// Clear removes all items from the queue.
func (q *Queue[T]) Clear() {
	q.list = nil
}
