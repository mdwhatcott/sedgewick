package ch1p5

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func weightedQuickFindAll(ints []int) Union {
	uf := NewWeightedQuickFind(ints[0])
	for x := 1; x < len(ints); x += 2 {
		uf.Union(ints[x], ints[x+1])
	}
	return uf
}

func TestWeightedQuickFindTiny(t *testing.T) {
	should.So(t, weightedQuickFindAll(tiny).Count(), should.Equal, 2)
}
func TestWeightedQuickFindMedium(t *testing.T) {
	should.So(t, weightedQuickFindAll(medium).Count(), should.Equal, 3)
}
func TestWeightedQuickFindLarge(t *testing.T) {
	should.So(t, weightedQuickFindAll(large).Count(), should.Equal, 6)
}
