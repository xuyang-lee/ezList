package list

func Reverse[T any](s []T) {
	last := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[last-i] = s[last-i], s[i]
	}
}

func Extend[T any](s []T, ext []T) []T {
	newList := make([]T, len(s)+len(ext))
	copy(newList, s)
	copy(newList[len(s):], ext)
	return newList
}

func Count[T comparable](s []T, e T) int {
	count := 0
	for _, v := range s {
		if v == e {
			count++
		}
	}
	return count
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func Distinct[T comparable](s []T) []T {
	m := make(map[T]struct{})
	for _, v := range s {
		m[v] = struct{}{}
	}
	newList := make([]T, 0, len(m))
	for k := range m {
		newList = append(newList, k)
	}
	return newList
}
