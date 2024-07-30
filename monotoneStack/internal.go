package monotoneStack

import "golang.org/x/exp/constraints"

type compareFunc[T any] func(top, elem T) bool

func compareGet[T constraints.Ordered](t MonotoneType, strict bool) compareFunc[T] {

	switch t {
	case MonotoneTypeIncrease:
		if strict {
			return increaseStrict[T]
		} else {
			return increase[T]
		}
	case MonotoneTypeDecrease:
		if strict {
			return decreaseStrict[T]
		} else {
			return decrease[T]
		}
	default:
		if strict {
			return increaseStrict[T]
		} else {
			return increase[T]
		}
	}
}

func increaseStrict[T constraints.Ordered](top, elem T) bool {
	return top < elem
}

func increase[T constraints.Ordered](top, elem T) bool {
	return top <= elem
}

func decreaseStrict[T constraints.Ordered](top, elem T) bool {
	return top > elem
}

func decrease[T constraints.Ordered](top, elem T) bool {
	return top >= elem
}
