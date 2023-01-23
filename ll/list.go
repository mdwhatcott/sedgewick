package ll

type l[T any] struct {
	Value T
	Next  *l[T]
}
type List[T any] struct {
	first *l[T]
	n     int
}

func NewList[T any]() *List[T]      { return new(List[T]) }
func (this *List[T]) IsEmpty() bool { return this.n == 0 }
func (this *List[T]) Len() int      { return this.n }
func (this *List[T]) Insert(n int, value T) {
	defer func() { this.n++ }()
	item := this.first
	for ; n > 0; n-- {
		item = item.Next
	}
	if item == nil {
		this.first = &l[T]{Value: value}
	} else {
		add := &l[T]{Value: value, Next: item.Next}
		item.Next = add
	}
}
func (this *List[T]) Pop(n int) T {
	defer func() { this.n-- }()
	if n == 0 {
		value := this.first.Value
		this.first = this.first.Next
		return value
	}
	item := this.first
	for n > 1 {
		item = item.Next
		n--
	}
	value := item.Next.Value
	item.Next = item.Next.Next
	return value
}
