package ch1p5

type UnionFind struct {
	ids   []int
	count int
}

func NewUnionFind(n int) *UnionFind {
	ints := make([]int, n)
	for x := 0; x < n; x++ {
		ints[x] = x
	}
	return &UnionFind{
		ids:   ints,
		count: n,
	}
}

func (this *UnionFind) Union(a, b int) {
	p := this.Find(a)
	q := this.Find(b)
	if p == q {
		return
	}

	for x := 0; x < len(this.ids); x++ {
		if this.Find(x) == p {
			this.ids[x] = this.Find(q)
		}
	}

	this.count--
}

func (this *UnionFind) Find(n int) int {
	return this.ids[n]
}

func (this *UnionFind) Connected(p, q int) bool {
	return this.Find(p) == this.Find(q)
}

func (this *UnionFind) Count() int {
	return this.count
}
