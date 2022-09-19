package ch1p4

// ThreeSumNaive https://algs4.cs.princeton.edu/14analysis/ThreeSum.java.html
func ThreeSumNaive(a []int) (count int) {
	for n, i := len(a), 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if a[i]+a[j]+a[k] == 0 {
					count++
				}
			}
		}
	}
	return count
}
