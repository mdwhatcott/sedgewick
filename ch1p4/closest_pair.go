package ch1p4

import "sort"

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// ClosestPairSlow executes in quadratic time
func ClosestPairSlow(all []int) int {
	a, b, min := 0, 0, 0xFFFFFFFF
	for x := 0; x < len(all); x++ {
		for y := x + 1; y < len(all); y++ {
			A, B := all[x], all[y]
			if distance := abs(A - B); distance < min {
				min, a, b = distance, A, B
			}
		}
	}
	return abs(a - b)
}

// ClosestPairFast executes in linearithmic time
func ClosestPairFast(all []int) int {
	sort.Ints(all)
	a, b, min := 0, 0, 0xFFFFFFFF
	for x := 0; x < len(all)-1; x++ {
		A, B := all[x], all[x+1]
		if distance := abs(A - B); distance < min {
			min, a, b = distance, A, B
		}
	}
	return abs(a - b)
}
