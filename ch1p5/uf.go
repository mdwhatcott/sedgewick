package ch1p5

type Union interface {
	Union(a, b int)
	Find(n int) int
	Connected(p, q int) bool
	Count() int
}
