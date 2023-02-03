package ch2p2

import (
	"testing"

	"github.com/mdwhatcott/testing/should"

	"sedgewick/ll"
)

func TestLLMergeSort(t *testing.T) {
	list := ll.NewList[int]()
	for x := 0; x < 10; x++ {
		list.Insert(x)
	}

	sorted := LLMergeSort(list.Head)

	should.So(t, ll.Slice(sorted), should.Equal, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}
