package ch1p4

import (
	"math/rand"
	"testing"
	"time"

	"github.com/mdwhatcott/go-collections/set"
	"github.com/mdwhatcott/testing/should"
)

func TestClosestPairSuite(t *testing.T) {
	should.Run(&ClosestPairSuite{T: should.New(t)}, should.Options.UnitTests())
}

type ClosestPairSuite struct {
	*should.T
}

func randomInts(n int) (results []int) {
	min, max := -1000000, 1000000

	unique := set.New[int](n)
	for len(unique) < n {
		i := rand.Intn(max-min) + min
		if i == 0 {
			continue
		}
		unique.Add(i)
	}
	return unique.Slice()
}

func (this *ClosestPairSuite) TestSlow() {
	ints := randomInts(10000)

	a := time.Now()
	slow := ClosestPairSlow(ints)
	this.Log("slow:", time.Since(a))

	b := time.Now()
	fast := ClosestPairFast(ints)
	this.Log("fast:", time.Since(b))

	this.So(slow, should.Equal, fast)
	this.Log(slow)
}
