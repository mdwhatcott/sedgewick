package ch1p3

import (
	"testing"

	"github.com/mdwhatcott/go-collections/set"
	"github.com/mdwhatcott/go-collections/stack"

	"github.com/mdwhatcott/testing/should"
)

// TestExercise1_3_4: Write a stack client that reads
// in text and uses a stack to determine whether its
// parentheses are properly balanced.
func TestExercise1_3_4(t *testing.T) {
	should.So(t, ParenthesesBalanced(""), should.BeTrue)
	should.So(t, ParenthesesBalanced("("), should.BeFalse)
	should.So(t, ParenthesesBalanced("(]"), should.BeFalse)
	should.So(t, ParenthesesBalanced("[)"), should.BeFalse)
	should.So(t, ParenthesesBalanced("{)"), should.BeFalse)
	should.So(t, ParenthesesBalanced("()"), should.BeTrue)
	should.So(t, ParenthesesBalanced("[]"), should.BeTrue)
	should.So(t, ParenthesesBalanced("{}"), should.BeTrue)
	should.So(t, ParenthesesBalanced("(())"), should.BeTrue)
	should.So(t, ParenthesesBalanced("(()]"), should.BeFalse)
	should.So(t, ParenthesesBalanced("[(])"), should.BeFalse)
	should.So(t, ParenthesesBalanced("[()]{}{[()()]()}"), should.BeTrue)
}

var (
	opens  = set.New[rune](0)
	closes = set.New[rune](0)
	pairs  = make(map[rune]rune)
)

func init() {
	for _, match := range []string{"()", "[]", "{}"} {
		opening := rune(match[0])
		closing := rune(match[1])
		opens.Add(opening)
		closes.Add(closing)
		pairs[closing] = opening
	}
}

func ParenthesesBalanced(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	items := stack.New[rune](0)
	for _, c := range s {
		if opens.Contains(c) {
			items.Push(c)
		} else if closes.Contains(c) {
			if items.Pop() != pairs[c] {
				return false
			}
		}
	}
	return items.Empty()
}
