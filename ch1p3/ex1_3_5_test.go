package ch1p3

import (
	"fmt"
	"testing"

	"github.com/mdwhatcott/testing/should"

	"sedgewick/ll"
)

func TestExercise1_3_5(t *testing.T) {
	should.So(t, Bin(0), should.Equal, "0")
	should.So(t, Bin(1), should.Equal, "1")
	should.So(t, Bin(2), should.Equal, "10")
	should.So(t, Bin(50), should.Equal, "110010")
}

func Bin(n int) (result string) {
	if n == 0 {
		return "0"
	}
	s := ll.NewStack[string]()
	for n > 0 {
		s.Push(fmt.Sprint(n % 2))
		n /= 2
	}
	for !s.IsEmpty() {
		result += s.Pop()
	}
	return result
}
