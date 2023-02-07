package ch2p3

import (
	"fmt"
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
	fmt.Println(a)
	i := lo
	j := hi + 1
	v := a[lo]
	for {
		for {
			i++
			if a[i] < v || i == hi {
				break
			}
		}
		for {
			j--
			if v < a[j] || j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		exchange(a, i, j)
		fmt.Println(a)
	}
	exchange(a, lo, j)
	fmt.Println(a)
	return j
}
func exchange[T generic.LessThan](a []T, i, j int) {
	a[i], a[j] = a[j], a[i]
}
