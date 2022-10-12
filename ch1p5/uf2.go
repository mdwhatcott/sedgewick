package ch1p5

type QuickUnion struct {
	ids   []int
	count int
}

func NewQuickUnion(n int) *QuickUnion {
	ints := make([]int, n)
	for x := 0; x < n; x++ {
		ints[x] = x
	}
	return &QuickUnion{
		ids:   ints,
		count: n,
	}
}

func (this *QuickUnion) Union(a, b int) {
	i := this.Find(a)
	j := this.Find(b)
	if i == j {
		return
	}
	this.ids[i] = j
	this.count--
}

func (this *QuickUnion) Find(p int) int {
	for p != this.ids[p] {
		p = this.ids[p]
	}
	return p
}

func (this *QuickUnion) Connected(p, q int) bool {
	return this.Find(p) == this.Find(q)
}

func (this *QuickUnion) Count() int {
	return this.count
}
