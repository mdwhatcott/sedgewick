package ch2p3

import (
	"math/rand"

	"sedgewick/util/generic"
)

func QuickSort[T generic.LessThan](a []T) {
	rand.Shuffle(len(a), func(i, j int) { exchange(a, i, j) })
	sort(a, 0, len(a)-1)
}
func sort[T generic.LessThan](a []T, lo, hi int) {
	if hi <= lo {
		return
	}
	j := partition(a, lo, hi)
	sort(a, lo, j-1)
	sort(a, j+1, hi)
}
func partition[T generic.LessThan](a []T, lo, hi int) int {
	i := lo
	j := hi + 1
	v := a[lo]
	for {
		for i++; a[i] < v && i != hi; i++ {
		}
		for j--; v < a[j] && j != lo; j-- {
		}
		if i >= j {
			break
		}
		exchange(a, i, j)
	}
	exchange(a, lo, j)
	return j
}
func exchange[T generic.LessThan](a []T, i, j int) {
	a[i], a[j] = a[j], a[i]
}
