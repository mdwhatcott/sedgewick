package ch2p2

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestTopDownMergeSort(t *testing.T) {
	items := []rune("MERGESORTEXAMPLE")
	TopDownMergeSort(items)
	should.So(t, string(items), should.Equal, "AEEEEGLMMOPRRSTX")
}
