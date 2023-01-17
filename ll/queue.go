package ll

type q[T any] struct {
	Value T
	Next  *q[T]
}

type Queue[T any] struct {
	first *q[T]
	last  *q[T]
	n     int
}

func NewQueue[T any]() *Queue[T]     { return new(Queue[T]) }
func (this *Queue[T]) IsEmpty() bool { return this.n == 0 }
func (this *Queue[T]) Len() int      { return this.n }
func (this *Queue[T]) Enqueue(value T) {
	oldLast := this.last
	this.last = new(q[T])
	this.last.Value = value
	this.last.Next = nil
	if this.IsEmpty() {
		this.first = this.last
	} else {
		oldLast.Next = this.last
	}
	this.n++
}
func (this *Queue[T]) Dequeue() T {
	value := this.first.Value
	this.first = this.first.Next
	this.n--
	if this.IsEmpty() {
		this.last = nil
	}
	return value
}
