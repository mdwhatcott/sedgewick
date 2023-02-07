package ch2p3

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestQuick(t *testing.T) {
	input := []rune("QUICKSORTEXAMPLE")
	QuickSort(input)
	should.So(t, string(input), should.Equal, "ACEEIKLMOPQRSTUX")

	other := []string{"K", "R", "A", "T", "E", "L", "E", "P", "U", "I", "M", "Q", "C", "X", "O", "S"}
	p := partition(other, 0, len(other)-1)
	t.Log(p)
	should.So(t, strings.Join(other, ""), should.Equal, "ECAIEKLPUTMQRXOS")
}
