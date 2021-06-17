package parser

import (
	"testing"
	"github.com/stretchr/testify/require"
)

var testDate = []struct {
	input    string
	expected  Node
	fail     bool
	errMsg   string
}{
	{
		input: "2021 10 1",
		expected: &DateTime{
			Year: 2021,
    		Month: 10,
    		Day: 1,
		},
	},
	{
		input: "2011 1 10",
		expected: &DateTime{
			Year: 2011,
    		Month: 1,
    		Day: 10,
		},
	},
	{
		input: "1970 01 01",
		expected: &DateTime{
			Year: 1970,
    		Month: 1,
    		Day: 1,
		},
	},
	{
		input: "1971 10 10",
		expected: &DateTime{
			Year: 1971,
    		Month: 10,
    		Day: 10,
		},
	},
	{
		input: "2010/10/10",
		expected: &DateTime{
			Year: 2010,
    		Month: 10,
    		Day: 10,
		},
	},
	{
		input: "2010-1-1",
		expected: &DateTime{
			Year: 2010,
    		Month: 1,
    		Day: 1,
		},
	},
	{
		input: "2021-Feb-1",
		expected: &DateTime{
			Year: 2021,
    		Month: 2,
    		Day: 1,
		},
	},
	{
		input: "oct 7, 2021",
		expected: &DateTime{
			Year: 2021,
    		Month: 10,
    		Day: 7,
		},
	},
	{
		input: "oct 7, 2021",
		expected: &DateTime{
			Year: 2021,
    		Month: 10,
    		Day: 7,
		},
	},
	{
		input: "Oct. 7, 2021",
		expected: &DateTime{
			Year: 2021,
    		Month: 10,
    		Day: 7,
		},
	},
	{
		input: "7 oct 1970",
		expected: &DateTime{
			Year: 1970,
    		Month: 10,
    		Day: 7,
		},
	},
	{
		input: "7 oct 70",
		expected: &DateTime{
			Year: 1970,
    		Month: 10,
    		Day: 7,
		},
	},
	{
		input: "2021年7月1日",
		expected: &DateTime{
			Year: 2021,
    		Month: 7,
    		Day: 1,
		},
	},
}

func TestParserDateTime(t *testing.T) {
	for i, test := range testDate {
		out, err := ParseDate(test.input)
		require.NotEqual(t, err, errUnexpected, "unexpected error occurred")

		require.Equal(t, test.expected, out, "%d: input %q", i, test.input)
	}
}
