package ch2p3

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestQuickMedian(t *testing.T) {
	input := []rune("QUICKSORTEXAMPLE")
	QuickSortMedian(input)
	should.So(t, string(input), should.Equal, "ACEEIKLMOPQRSTUX")
}

func TestMedian3Index(t *testing.T) {
	should.So(t, median3Index([]int{24, 19, 199}, 0, 1, 2), should.Equal, 0)
	should.So(t, median3Index([]int{1, 19, 199}, 0, 1, 2), should.Equal, 1)
	should.So(t, median3Index([]int{1000, 19, 199}, 0, 1, 2), should.Equal, 2)
}
