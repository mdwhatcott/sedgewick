package ll

type Node[T any] struct {
	Value T
	Next  *Node[T]
}
type List[T any] struct {
	Head *Node[T]
	n    int
}

func NewList[T any]() *List[T]      { return new(List[T]) }
func (this *List[T]) IsEmpty() bool { return this.n == 0 }
func (this *List[T]) Len() int      { return this.n }
func (this *List[T]) Insert(value T) {
	this.Head = &Node[T]{Value: value, Next: this.Head}
	this.n++
}
func (this *List[T]) Slice() (result []T) {
	return Slice(this.Head)
}
func Slice[T any](head *Node[T]) (result []T) {
	item := head
	for item != nil {
		result = append(result, item.Value)
		item = item.Next
	}
	return result
}
