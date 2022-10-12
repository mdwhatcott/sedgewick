package ch1p5

type WeightedQuickFind struct {
	ids   []int
	count int
	sizes []int
}

func NewWeightedQuickFind(n int) *WeightedQuickFind {
	ints := make([]int, n)
	for x := 0; x < n; x++ {
		ints[x] = x
	}
	sizes := make([]int, n)
	for x := 0; x < n; x++ {
		sizes[x] = 1
	}
	return &WeightedQuickFind{
		ids:   ints,
		count: n,
		sizes: sizes,
	}
}

func (this *WeightedQuickFind) Union(a, b int) {
	i := this.Find(a)
	j := this.Find(b)
	if i == j {
		return
	}
	if this.sizes[i] < this.sizes[j] {
		this.ids[i] = j
		this.sizes[j] += this.sizes[i]
	} else {
		this.ids[j] = i
		this.sizes[i] += this.sizes[j]
	}
	this.count--
}

func (this *WeightedQuickFind) Find(p int) int {
	for p != this.ids[p] {
		p = this.ids[p]
	}
	return p
}

func (this *WeightedQuickFind) Connected(p, q int) bool {
	return this.Find(p) == this.Find(q)
}

func (this *WeightedQuickFind) Count() int {
	return this.count
}
