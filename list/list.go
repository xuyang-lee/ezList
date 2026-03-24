package list

import (
	"github.com/xuyang-lee/ezSet/orderSet"
	"github.com/xuyang-lee/ezSet/set"
	"math/rand"
	"time"
)

// Reverse slice s
func Reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
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
	for i := range s {
		if !exc.Contains(s[i]) {
			newList = append(newList, s[i])
		}
	}
	return newList

}

// Count the number of e in s
func Count[T comparable](s []T, e T) int {
	count := 0
	for i := range s {
		if s[i] == e {
			count++
		}
	}
	return count
}

// Contains checks if e is in s
func Contains[T comparable](s []T, e T) bool {
	for i := range s {
		if s[i] == e {
			return true
		}
	}
	return false
}

// Distinct returns a new slice with duplicates removed, but do not guarantee the order of results
func Distinct[T comparable](s []T) []T {
	return set.NewSetWithSlice(s).List()
}

// OrderDistinct returns a new slice with duplicates removed, The result keeps the order of the first appearance of the elements
func OrderDistinct[T comparable](s []T) []T {
	return orderSet.NewOrderSetWithSlice(s).List()
}

// Filter returns a new slice with elements that satisfy the predicate 'keep'
func Filter[T any](s []T, keep func(T) bool) []T {
	var newList []T
	for i := range s {
		if keep(s[i]) {
			newList = append(newList, s[i])
		}
	}
	return newList
}

// ProcessEach applies the function f to each element in the slice s
func ProcessEach[T any](s []T, process func(T) T) {
	for i := range s {
		s[i] = process(s[i])
	}
}

// Any checks if there is any element in the slice `s` that is not equal to the zero value of the type `T`
func Any[T comparable](s []T) bool {
	var e T
	for i := range s {
		if s[i] != e {
			return true
		}
	}
	return false
}

// All checks if all elements in the slice `s` are non-zero values for the type `T`
func All[T comparable](s []T) bool {
	var e T
	for i := range s {
		if s[i] == e {
			return false
		}
	}
	return true
}

// Overlap returns overlap between a and b,
// if there isn't any overlap between a and b, return empty slice of type T and false
func Overlap[T comparable](a, b []T) ([]T, bool) {

	s := set.NewSetWithSlice(a).Intersect(set.NewSetWithSlice(b))
	return s.List(), s.Len() > 0
}

// IndexOf returns the index of the first occurrence of t in s, or -1 if t is not present in s
func IndexOf[T comparable](s []T, t T) int {
	for i := range s {
		if s[i] == t {
			return i
		}
	}
	return -1
}

// Shuffle the slice a
func Shuffle[T any](a []T) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := len(a) - 1; i > 0; i-- {
		j := r.Intn(i + 1)      // 生成一个0到i的随机索引
		a[i], a[j] = a[j], a[i] // 交换两个索引处的元素
	}
}

// Extract info from Slice  of type T to slice of type R
func Extract[T any, R any](s []T, transform func(T) R) []R {
	newList := make([]R, len(s))
	for i := range s {
		newList[i] = transform(s[i])
	}
	return newList
}

// ExtractPtr info from Slice  of type T to slice of type R
func ExtractPtr[T any, R any](s []T, transform func(*T) R) []R {
	newList := make([]R, len(s))
	for i := range s {
		newList[i] = transform(&s[i])
	}
	return newList
}

// ToMap converts a slice to a map
//
// keyFn can get key of map from elem
func ToMap[T any, R comparable](s []T, keyFn func(T) R) map[R]T {
	m := make(map[R]T)
	for i := range s {
		m[keyFn(s[i])] = s[i]
	}
	return m
}

func ToMapGroup[T any, R comparable](s []T, keyFn func(T) R) map[R][]T {
	m := make(map[R][]T)
	for i := range s {
		k := keyFn(s[i])
		m[k] = append(m[k], s[i])
	}
	return m
}

// Paginate paginate slice
func Paginate[T any](slice []T, page, size int) []T {

	if page <= 0 || size <= 0 {
		panic("page and size must be positive")
	}

	offset := (page - 1) * size
	endIndex := offset + size

	// deal start over index
	if offset >= len(slice) {
		return []T{}
	}

	// deal end over index
	if endIndex > len(slice) {
		endIndex = len(slice)
	}

	return slice[offset:endIndex]
}

// Random to get a rand elem from slice
func Random[T any](slice []T) T {
	source := rand.NewSource(time.Now().UnixNano())
	i := rand.New(source).Intn(len(slice))
	return slice[i]
}

func Pointers[T any](r []T) []*T {
	res := make([]*T, len(r))
	for i := range r {
		res[i] = &r[i]
	}
	return res
}

func Values[T any](s []*T) []T {
	res := make([]T, len(s))
	for i := range s {
		if s[i] != nil {
			res[i] = *s[i]
		}
	}
	return res
}
