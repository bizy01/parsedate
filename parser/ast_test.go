package parser

import (
	"testing"
)


func TestYear(t *testing.T) {
	var tokens = []Token{
		Token{
			Typ: DIGIT,
			Pos: 0,
			Val: "2",
		},
		Token{
			Typ: DIGIT,
			Pos: 1,
			Val: "0",
		},
		Token{
			Typ: DIGIT,
			Pos: 2,
			Val: "2",
		},
		Token{
			Typ: DIGIT,
			Pos: 3,
			Val: "1",
		},
	}

	year := JoinYear(tokens...)
	t.Log("year ====>", year)
}