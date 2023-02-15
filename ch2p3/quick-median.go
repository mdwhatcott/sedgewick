package ch2p3

import (
	"math/rand"

	"sedgewick/util/generic"
)

func QuickSortMedian[T generic.LessThan](a []T) {
	rand.Shuffle(len(a), func(i, j int) { exchange(a, i, j) })
	quickSortMedian(a, 0, len(a)-1)
}

const CUTOFF = 8

func quickSortMedian[T generic.LessThan](v []T, lo, hi int) {
	n := hi - lo + 1
	if n <= CUTOFF {
		insertionSort(v, lo, hi)
		return
	}

	m := median3Index(v, lo, lo+n/2, hi)
	exchange(v, m, lo)

	j := partition(v, lo, hi)
	quickSortMedian(v, lo, j-1)
	quickSortMedian(v, j+1, hi)
}

func insertionSort[T generic.LessThan](v []T, lo, hi int) {
	for i := lo; i <= hi; i++ {
		for j := i; j > lo && v[j] < v[j-1]; j-- {
			exchange(v, j, j-1)
		}
	}
}

func median3Index[T generic.LessThan](v []T, i, j, k int) int {
	if v[i] > v[j] {
		i, j = j, i
	}
	if v[j] > v[k] {
		j, k = k, j
	}
	return j
}
