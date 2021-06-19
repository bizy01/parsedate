package parser

import (
	"testing"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	input      string
	expected   []Token
	fail       bool
}

var tests = []struct {
	name  string
	tests []testCase
}{
	{
		name: "digit",
		tests: []testCase{
			{
				input:    "2",
				expected: []Token{
					{DIGIT, 0, "2"},
				},
			},
			{
				input:    "-",
				expected: []Token{
					{DASH, 0, "-"},
				},
			},
			{
				input:    "/",
				expected: []Token{
					{DIV, 0, "/"},
				},
			},
			{
				input:    ".",
				expected: []Token{
					{DOT, 0, "."},
				},
			},
			{
				input:    ":",
				expected: []Token{
					{COLON, 0, ":"},
				},
			},
			{
				input:    ",",
				expected: []Token{
					{COMMA, 0, ","},
				},
			},
			{
				input:    " ",
				expected: []Token{
					{SPACE, 0, " "},
				},
			},
            {
				input: "2=",
				fail: true,
			},
			{
				input:    "2020",
				expected: []Token{
					{DIGIT, 0, "2"},
					{DIGIT, 1, "0"},
					{DIGIT, 2, "2"},
					{DIGIT, 3, "0"},
				},
			},
			{
				input:    "2020-1-1",
				expected: []Token{
					{DIGIT, 0, "2"},
					{DIGIT, 1, "0"},
					{DIGIT, 2, "2"},
					{DIGIT, 3, "0"},
					{DASH,  4, "-"},
					{DIGIT, 5, "1"},
					{DASH, 6, "-"},
					{DIGIT, 7, "1"},
				},
			},
			{
				input:    "2020-01-01 10:10:10",
				expected: []Token{
					{DIGIT, 0, "2"},
					{DIGIT, 1, "0"},
					{DIGIT, 2, "2"},
					{DIGIT, 3, "0"},
					{DASH,  4, "-"},
					{DIGIT, 5, "0"},
					{DIGIT, 6, "1"},
					{DASH,  7, "-"},
					{DIGIT, 8, "0"},
					{DIGIT, 9, "1"},
					{SPACE, 10, " "},
					{DIGIT, 11, "1"},
					{DIGIT, 12, "0"},
					{COLON, 13, ":"},
					{DIGIT, 14, "1"},
					{DIGIT, 15, "0"},
					{COLON, 16, ":"},
					{DIGIT, 17, "1"},
					{DIGIT, 18, "0"},
				},
			},
			{
				input:    "2020-01-01 10:10:10.1",
				expected: []Token{
					{DIGIT, 0, "2"},
					{DIGIT, 1, "0"},
					{DIGIT, 2, "2"},
					{DIGIT, 3, "0"},
					{DASH,  4, "-"},
					{DIGIT, 5, "0"},
					{DIGIT, 6, "1"},
					{DASH,  7, "-"},
					{DIGIT, 8, "0"},
					{DIGIT, 9, "1"},
					{SPACE, 10, " "},
					{DIGIT, 11, "1"},
					{DIGIT, 12, "0"},
					{COLON, 13, ":"},
					{DIGIT, 14, "1"},
					{DIGIT, 15, "0"},
					{COLON, 16, ":"},
					{DIGIT, 17, "1"},
					{DIGIT, 18, "0"},
					{DOT,   19, "."},
					{DIGIT, 20, "1"},
				},
			},
			{
				input:    "2020-01-01 10:10:10.123",
				expected: []Token{
					{DIGIT, 0, "2"},
					{DIGIT, 1, "0"},
					{DIGIT, 2, "2"},
					{DIGIT, 3, "0"},
					{DASH,  4, "-"},
					{DIGIT, 5, "0"},
					{DIGIT, 6, "1"},
					{DASH,  7, "-"},
					{DIGIT, 8, "0"},
					{DIGIT, 9, "1"},
					{SPACE, 10, " "},
					{DIGIT, 11, "1"},
					{DIGIT, 12, "0"},
					{COLON, 13, ":"},
					{DIGIT, 14, "1"},
					{DIGIT, 15, "0"},
					{COLON, 16, ":"},
					{DIGIT, 17, "1"},
					{DIGIT, 18, "0"},
					{DOT,   19, "."},
					{DIGIT, 20, "1"},
					{DIGIT, 21, "2"},
					{DIGIT, 22, "3"},
				},
			},
			{
				input:    "2020-01-01 10:10:10.123456",
				expected: []Token{
					{DIGIT, 0, "2"},
					{DIGIT, 1, "0"},
					{DIGIT, 2, "2"},
					{DIGIT, 3, "0"},
					{DASH,  4, "-"},
					{DIGIT, 5, "0"},
					{DIGIT, 6, "1"},
					{DASH,  7, "-"},
					{DIGIT, 8, "0"},
					{DIGIT, 9, "1"},
					{SPACE, 10, " "},
					{DIGIT, 11, "1"},
					{DIGIT, 12, "0"},
					{COLON, 13, ":"},
					{DIGIT, 14, "1"},
					{DIGIT, 15, "0"},
					{COLON, 16, ":"},
					{DIGIT, 17, "1"},
					{DIGIT, 18, "0"},
					{DOT,   19, "."},
					{DIGIT, 20, "1"},
					{DIGIT, 21, "2"},
					{DIGIT, 22, "3"},
					{DIGIT, 23, "4"},
					{DIGIT, 24, "5"},
					{DIGIT, 25, "6"},
				},
			},
		},
	},
	{
		name: "keywords",
		tests: []testCase{
			{
				input:    "January",
				expected: []Token{{MONTH, 0, "January"}},
			},
			{
				input:    "February",
				expected: []Token{{MONTH, 0, "February"}},
			},
			{
				input:    "March",
				expected: []Token{{MONTH, 0, "March"}},
			},
			{
				input:    "April",
				expected: []Token{{MONTH, 0, "April"}},
			},
			{
				input:    "May",
				expected: []Token{{MONTH, 0, "May"}},
			},
			{
				input:    "June",
				expected: []Token{{MONTH, 0, "June"}},
			},
			{
				input:    "July",
				expected: []Token{{MONTH, 0, "July"}},
			},
			{
				input:    "August",
				expected: []Token{{MONTH, 0, "August"}},
			},
			{
				input:    "September",
				expected: []Token{{MONTH, 0, "September"}},
			},
			{
				input:    "October",
				expected: []Token{{MONTH, 0, "October"}},
			},
			{
				input:    "November",
				expected: []Token{{MONTH, 0, "November"}},
			},
			{
				input:    "December",
				expected: []Token{{MONTH, 0, "December"}},
			},
			{
				input:    "Jan",
				expected: []Token{{MONTH, 0, "Jan"}},
			},
			{
				input:    "Feb",
				expected: []Token{{MONTH, 0, "Feb"}},
			},
			{
				input:    "Mar",
				expected: []Token{{MONTH, 0, "Mar"}},
			},
			{
				input:    "Apr",
				expected: []Token{{MONTH, 0, "Apr"}},
			},
			{
				input:    "Jun",
				expected: []Token{{MONTH, 0, "Jun"}},
			},
			{
				input:    "Jul",
				expected: []Token{{MONTH, 0, "Jul"}},
			},
			{
				input:    "Aug",
				expected: []Token{{MONTH, 0, "Aug"}},
			},
			{
				input:    "Sept",
				expected: []Token{{MONTH, 0, "Sept"}},
			},
			{
				input:    "Oct",
				expected: []Token{{MONTH, 0, "Oct"}},
			},
			{
				input:    "Nov",
				expected: []Token{{MONTH, 0, "Nov"}},
			},
			{
				input:    "Dec",
				expected: []Token{{MONTH, 0, "Dec"}},
			},
			{
				input:    "oct 7, 1970",
				expected: []Token{
					{MONTH, 0, "oct"},
					{SPACE, 3, " "},
					{DIGIT, 4, "7"},
					{COMMA, 5, ","},
					{SPACE, 6, " "},
					{DIGIT, 7, "1"},
					{DIGIT, 8, "9"},
					{DIGIT, 9, "7"},
					{DIGIT, 10, "0"},
				},
			},
		},
	},
	{
		name: "error",
		tests: []testCase{
			{
				input:    "#",
				expected: []Token{{ERROR, 0, "unexpected character:#"}},
				fail: true,
			},
			{
				input:    "2012#",
				expected: []Token{{ERROR, 0, "unexpected character:#"}},
				fail: true,
			},
			{
				input:    "aaaa",
				expected: []Token{{ERROR, 0, "unexpected key word:aaaa"}},
				fail: true,
			},
		},
	},
}

func TestLexer(t *testing.T) {
	for _, typ := range tests {
		t.Run(typ.name, func(t *testing.T) {
			for i, test := range typ.tests {
				l := &Lexer{
					input: test.input,
					state: lexStatements,
				}

				var out []Token
				for l.state != nil {
					out = append(out, Token{})

					l.NextItem(&out[len(out)-1])
				}

                lastItem := out[len(out)-1]
				if test.fail {
					hasError := false
					for _, item := range out {
						if item.Typ == ERROR {
							hasError = true
						}
					}

					if !hasError {
						t.Logf("%d: input %q", i, test.input)
						t.Fatalf("expected lexing error but did not fail")
					}
					continue
				}

				if lastItem.Typ == ERROR {
					t.Logf("%d: input %q", i, test.input)
					t.Fatalf("unexpected lexing error at position %d: %v", lastItem.Pos, lastItem)
				}

				eofItem := Token{EOF, Pos(len(test.input)), ""}
				require.Equal(t, lastItem, eofItem, "%d: input %q", i, test.input)

				out = out[:len(out)-1]
				require.Equal(t, test.expected, out, "%d: input %q", i, test.input)
			}
		})
	}
}
