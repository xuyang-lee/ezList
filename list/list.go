package list

import "github.com/xuyang-lee/ezSet/set"

// Reverse slice s
func Reverse[T any](s []T) {
	last := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[last-i] = s[last-i], s[i]
	}
}

// Extend s with ext
func Extend[T any](s []T, ext []T) []T {
	newList := make([]T, len(s)+len(ext))
	copy(newList, s)
	copy(newList[len(s):], ext)
	return newList
}

// Exclude e from s
func Exclude[T comparable](s []T, e []T) []T {
	exc := set.NewSetWithSlice(e)

	var newList []T
	for _, v := range s {
		if !exc.Contains(v) {
			newList = append(newList, v)
		}
	}
	return newList

}

// Count the number of e in s
func Count[T comparable](s []T, e T) int {
	count := 0
	for _, v := range s {
		if v == e {
			count++
		}
	}
	return count
}

// Contains checks if e is in s
func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// Distinct returns a new slice with duplicates removed, but do not guarantee the order of results
func Distinct[T comparable](s []T) []T {
	return set.NewSetWithSlice(s).List()
}

// Filter returns a new slice with elements that satisfy the predicate f
func Filter[T any](s []T, f func(T) bool) []T {
	var newList []T
	for _, v := range s {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// Any checks if there is any element in the slice `s` that is not equal to the zero value of the type `T`
func Any[T comparable](s []T) bool {
	var e T
	for _, v := range s {
		if v != e {
			return true
		}
	}
	return false
}

// All checks if all elements in the slice `s` are non-zero values for the type `T`
func All[T comparable](s []T) bool {
	var e T
	for _, v := range s {
		if v == e {
			return false
		}
	}
	return true
}

// Overlap returns overlap between a and b,
// if there isn't any overlap between a and b, return empty slice of type T and false
func Overlap[T comparable](a, b []T) ([]T, bool) {

	s := set.NewSetWithSlice(a).Union(set.NewSetWithSlice(b))
	return s.List(), s.Len() > 0
}
