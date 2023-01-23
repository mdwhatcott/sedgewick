package ll

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestList(t *testing.T) {
	l := NewList[int]()
	l.Insert(0, 1)
	l.Insert(0, 2)
	l.Insert(0, 3)
	l.Insert(0, 4)
	l.Insert(0, 5)
	t.Log(Slice(l))
	should.So(t, l.Len(), should.Equal, 3)
	should.So(t, l.Pop(0), should.Equal, 3)
	should.So(t, l.Pop(0), should.Equal, 2)
	should.So(t, l.Pop(0), should.Equal, 1)
	should.So(t, l.IsEmpty(), should.BeTrue)
}

func Slice[T any](l *List[T]) (result []T) {
	for !l.IsEmpty() {
		result = append(result, l.Pop(0))
	}
	return result
}
