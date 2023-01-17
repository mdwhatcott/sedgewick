package ll

type s[T any] struct {
	Value T
	Next  *s[T]
}

type Stack[T any] struct {
	first *s[T]
	n     int
}

func NewStack[T any]() *Stack[T]     { return new(Stack[T]) }
func (this *Stack[T]) IsEmpty() bool { return this.n == 0 }
func (this *Stack[T]) Len() int      { return this.n }
func (this *Stack[T]) Push(value T) {
	oldFirst := this.first
	first := new(s[T])
	first.Value = value
	first.Next = oldFirst
	this.n++
}
func (this *Stack[T]) Pop() T {
	item := this.first.Value
	this.first = this.first.Next
	this.n--
	return item
}
