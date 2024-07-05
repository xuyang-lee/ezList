package stack

import (
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	a := []int{1, 2, 3}
	fmt.Printf("len:%v;cap:%v\n", len(a), cap(a))
	b := make([]int, len(a)+5)
	fmt.Printf("len:%v;cap:%v\n", len(b), cap(b))
	copy(b, a)
	fmt.Printf("len:%v;cap:%v\n", len(b), cap(b))

	c := []int{1, 2, 3}
	d := []int{4, 5, 6, 1, 2, 3}
	fmt.Printf("len:%v;cap:%v\n", len(c), cap(c))
	copy(c, d)
	fmt.Printf("c:%v\n", c)
}
