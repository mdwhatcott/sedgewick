package ll

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestList(t *testing.T) {
	l := NewList[int]()
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Insert(4)
	l.Insert(5)
	should.So(t, l.Len(), should.Equal, 5)
	should.So(t, l.Slice(), should.Equal, []int{5, 4, 3, 2, 1})
}
