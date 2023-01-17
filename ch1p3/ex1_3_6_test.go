package ch1p3

import (
	"testing"

	"github.com/mdwhatcott/go-collections/stack"

	"sedgewick/ll"
)

func TestExercise1_3_6(t *testing.T) {
	q := ll.NewQueue[int]()
	for x := 0; x < 10; x++ {
		q.Enqueue(x)
	}
	Reverse(q)
	LogAll(t, q)
}

func Reverse[T comparable](q *ll.Queue[T]) {
	s := stack.New[T](0)
	for !q.IsEmpty() {
		s.Push(q.Dequeue())
	}
	for !s.Empty() {
		q.Enqueue(s.Pop())
	}
}

func LogAll[T comparable](t *testing.T, q *ll.Queue[T]) {
	for !q.IsEmpty() {
		t.Log(q.Dequeue())
	}
}
