package ch1p5

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func quickUnionAll(ints []int) Union {
	uf := NewQuickUnion(ints[0])
	for x := 1; x < len(ints); x += 2 {
		uf.Union(ints[x], ints[x+1])
	}
	return uf
}

func TestQuickFindTiny(t *testing.T) {
	should.So(t, quickUnionAll(tiny).Count(), should.Equal, 2)
}
func TestQuickFindMedium(t *testing.T) {
	should.So(t, quickUnionAll(medium).Count(), should.Equal, 3)
}
func TestQuickFindLarge(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test in short mode.")
	}
	should.So(t, quickUnionAll(large).Count(), should.BeGreaterThan, 0)
}
