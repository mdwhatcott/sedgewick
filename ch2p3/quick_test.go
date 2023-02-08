package ch2p3

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestQuick(t *testing.T) {
	input := []rune("QUICKSORTEXAMPLE")
	QuickSort(input)
	should.So(t, string(input), should.Equal, "ACEEIKLMOPQRSTUX")
}
