package ch2p2

import (
	"fmt"
	"runtime/debug"
	"strings"
)

func TopDownMergeSort[T LessThan](a []T) {
	sortTopDown(a, make([]T, len(a)), 0, len(a)-1)
	fmt.Printf("          %#v\n", a)
}

func sortTopDown[T LessThan](a, aux []T, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	sortTopDown(a, aux, lo, mid)
	sortTopDown(a, aux, mid+1, hi)
	merge(a, aux, lo, mid, hi)
}

func merge[T LessThan](a, aux []T, lo, mid, hi int) {
	depth := len(strings.Split(string(debug.Stack()), "\n"))
	fmt.Printf("%03d merge(%#v, %#v, %#v, %#v, %#v)\n", depth, a, aux, lo, mid, hi)
	L, R := lo, mid+1

	for k := lo; k <= hi; k++ {
		aux[k] = a[k]
	}

	for k := lo; k <= hi; k++ {
		if L > mid {
			a[k] = aux[R]
			R++
		} else if R > hi {
			a[k] = aux[L]
			L++
		} else if aux[R] < aux[L] {
			a[k] = aux[R]
			R++
		} else {
			a[k] = aux[L]
			L++
		}
	}
}

type (
	Number interface {
		Integer | uintptr | float64 | float32
	}
	Integer interface {
		int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
	}
	LessThan interface {
		Number | string
	}
)
