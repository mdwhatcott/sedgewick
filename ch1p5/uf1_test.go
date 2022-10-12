package ch1p5

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func unionFindAll(ints []int) Union {
	uf := NewUnionFind(ints[0])
	for x := 1; x < len(ints); x += 2 {
		uf.Union(ints[x], ints[x+1])
	}
	return uf
}

func TestUnionFindTiny(t *testing.T) {
	should.So(t, unionFindAll(tiny).Count(), should.Equal, 2)
}
func TestUnionFindMedium(t *testing.T) {
	should.So(t, unionFindAll(medium).Count(), should.Equal, 3)
}
func TestUnionFindLarge(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test in short mode.")
	}
	should.So(t, unionFindAll(large).Count(), should.BeGreaterThan, 0)
}
