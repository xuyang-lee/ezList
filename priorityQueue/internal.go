package priorityQueue

import "golang.org/x/exp/constraints"

// compareFunc[T] func(src, des T) bool
//
// when src has a higher priority than des, returns true
type compareFunc[T any] func(src, des T) bool

func compareGet[T constraints.Ordered](t PriorityType) compareFunc[T] {

	switch t {
	case PriorityTypeGreater:
		return greater[T]
	case PriorityTypeLower:
		return lower[T]
	default:
		return greater[T]
	}
}

func greater[T constraints.Ordered](src, des T) bool {
	return src > des
}

func lower[T constraints.Ordered](src, des T) bool {
	return src < des
}
