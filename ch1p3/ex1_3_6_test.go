package ch1p3

import (
	"testing"

	"github.com/mdwhatcott/go-collections/queue"
	"github.com/mdwhatcott/go-collections/stack"
)

func TestExercise1_3_6(t *testing.T) {
	q := queue.New[int](0)
	for x := 0; x < 10; x++ {
		q.Enqueue(x)
	}
	t.Log(q.Slice())
	Reverse(q)
	t.Log(q.Slice())
}

func Reverse[T comparable](q *queue.Queue[T]) {
	s := stack.New[T](0)
	for !q.Empty() {
		s.Push(q.Dequeue())
	}
	for !s.Empty() {
		q.Enqueue(s.Pop())
	}
}
