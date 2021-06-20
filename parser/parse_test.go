package parser

import (
	"testing"
	"github.com/stretchr/testify/require"
)

var testDate = []struct {
	desc     string
	input    string
	expected  Node
	fail     bool
	errMsg   string
}{
	{
		desc: "ANSIC (Mon Jan _2 15:04:05 2006)",
		input: "Mon Jan  2 15:04:05 2021",
		expected: &DateTime{
			Year: 2021,
    		Month: 1,
    		Day: 2,
    		Hour: 15,
    		Minute: 4,
    		Second: 5,
		},
	},
	{
		desc: "ANSIC (Mon Jan 12 15:04:05 2006)",
		input: "Mon Jan 12 15:04:05 2021",
		expected: &DateTime{
			Year: 2021,
    		Month: 1,
    		Day: 12,
    		Hour: 15,
    		Minute: 4,
    		Second: 5,
		},
	},
	{
		desc: "UnixDate (Mon Jan _2 15:04:05 MST 2006)",
		input: "Mon Jan  2 15:04:05 MST 2021",
		expected: &DateTime{
			Year: 2021,
    		Month: 1,
    		Day: 2,
    		Hour: 15,
    		Minute: 4,
    		Second: 5,
		},
	},
	{
		desc: "UnixDate (Mon Jan 12 15:04:05 MST 2006)",
		input: "Mon Jan 12 15:04:05 MST 2021",
		expected: &DateTime{
			Year: 2021,
    		Month: 1,
    		Day: 12,
    		Hour: 15,
    		Minute: 4,
    		Second: 5,
		},
	},
	{
		desc: "RubyDate (Mon Jan 02 15:04:05 -0700 2006)",
		input: "Mon Jan 12 15:04:05 -0700 2021",
		expected: &DateTime{
			Year: 2021,
    		Month: 1,
    		Day: 12,
    		Hour: 15,
    		Minute: 4,
    		Second: 5,
		},
	},
	{
		desc: "RFC822 (02 Jan 06 15:04 MST)",
		input: "02 Jan 21 15:04 MST",
		expected: &DateTime{
			Year: 2021,
    		Month: 1,
    		Day: 2,
    		Hour: 15,
    		Minute: 4,
		},
	},
	{
		desc: "RFC822 (02 Jan 06 15:04 MST)",
		input: "02 Jan 21 15:04:05 MST",
		expected: &DateTime{
			Year: 2021,
    		Month: 1,
    		Day: 2,
    		Hour: 15,
    		Minute: 4,
    		Second: 5,
		},
	},
	{
		desc: "RFC822Z (02 Jan 06 15:04 -0700)",
		input: "02 Jan 21 15:04 -0700",
		expected: &DateTime{
			Year: 2021,
    		Month: 1,
    		Day: 2,
    		Hour: 15,
    		Minute: 4,
		},
	},
	{
		desc: "RFC850 (Monday, 02-Jan-06 15:04:05 MST)",
		input: "Monday, 02-Jan-21 15:04:05 MST",
		expected: &DateTime{
			Year: 2021,
    		Month: 1,
    		Day: 2,
    		Hour: 15,
    		Minute: 4,
    		Second: 5,
		},
	},
	{
		desc: "RFC3339 (2006-01-02T15:04:05Z07:00)",
		input: "2021-01-02T15:04:05Z07:00",
		expected: &DateTime{
			Year: 2021,
    		Month: 1,
    		Day: 2,
    		Hour: 15,
    		Minute: 4,
    		Second: 5,
		},
	},
	{
		desc: "RFC3339 (2006-01-02T15:04:05.999999999Z07:00)",
		input: "2021-01-02T15:04:05.999999999Z07:00",
		expected: &DateTime{
			Year: 2021,
    		Month: 1,
    		Day: 2,
    		Hour: 15,
    		Minute: 4,
    		Second: 5,
    		Millisecond: 999999999,
		},
	},
	// {
	// 	input: "Mon, 02 Jan 2021 15:04:05 MST",
	// 	expected: &DateTime{
	// 		Year: 2021,
 //    		Month: 1,
 //    		Day: 2,
 //    		Hour: 15,
 //    		Minute: 4,
 //    		Second: 5,
	// 	},
	// },
	{
		input: "Mon, 02 Jan 2021 15:04:05 -0700",
		expected: &DateTime{
			Year: 2021,
    		Month: 1,
    		Day: 2,
    		Hour: 15,
    		Minute: 4,
    		Second: 5,
		},
	},
	// {
	// 	input: "2011 1 10",
	// 	expected: &DateTime{
	// 		Year: 2011,
 //    		Month: 1,
 //    		Day: 10,
	// 	},
	// },
	// {
	// 	input: "1970 01 01",
	// 	expected: &DateTime{
	// 		Year: 1970,
 //    		Month: 1,
 //    		Day: 1,
	// 	},
	// },
	// {
	// 	input: "1971 10 10",
	// 	expected: &DateTime{
	// 		Year: 1971,
 //    		Month: 10,
 //    		Day: 10,
	// 	},
	// },
	// {
	// 	input: "2010/10/10",
	// 	expected: &DateTime{
	// 		Year: 2010,
 //    		Month: 10,
 //    		Day: 10,
	// 	},
	// },
	// {
	// 	input: "2010-1-1",
	// 	expected: &DateTime{
	// 		Year: 2010,
 //    		Month: 1,
 //    		Day: 1,
	// 	},
	// },
	// {
	// 	input: "2021-Feb-1",
	// 	expected: &DateTime{
	// 		Year: 2021,
 //    		Month: 2,
 //    		Day: 1,
	// 	},
	// },
	// {
	// 	input: "oct 7, 2021",
	// 	expected: &DateTime{
	// 		Year: 2021,
 //    		Month: 10,
 //    		Day: 7,
	// 	},
	// },
	// {
	// 	input: "oct 7, 2021",
	// 	expected: &DateTime{
	// 		Year: 2021,
 //    		Month: 10,
 //    		Day: 7,
	// 	},
	// },
	// {
	// 	input: "Oct. 7, 2021",
	// 	expected: &DateTime{
	// 		Year: 2021,
 //    		Month: 10,
 //    		Day: 7,
	// 	},
	// },
	// {
	// 	input: "7 oct 1970",
	// 	expected: &DateTime{
	// 		Year: 1970,
 //    		Month: 10,
 //    		Day: 7,
	// 	},
	// },
	// {
	// 	input: "7 oct 70",
	// 	expected: &DateTime{
	// 		Year: 1970,
 //    		Month: 10,
 //    		Day: 7,
	// 	},
	// },
	// {
	// 	input: "2021年7月1日",
	// 	expected: &DateTime{
	// 		Year: 2021,
 //    		Month: 7,
 //    		Day: 1,
	// 	},
	// },
	// {
	// 	input: "2020-01-01 10:10:10.123456",
	// 	expected: &DateTime{
	// 		Year: 2021,
 //    		Month: 1,
 //    		Day: 1,
 //    		Hour: 10,
 //    		Minute: 10,
 //    		Second: 10,
 //    		Millisecond: 123456,
	// 	},
	// 	fail: true,
	// },
	// {
	// 	input: "2021",
	// 	expected: &DateTime{
	// 		Year: 2021,
 //    		Month: 0,
 //    		Day: 0,
	// 	},
	// 	fail: true,
	// },
}

func TestParserDateTime(t *testing.T) {
	for i, test := range testDate {
		out, err := ParseDate(test.input)

		if test.fail {
			hasError := false

			if err != nil {
				hasError = true
			}

			if !hasError {
				t.Logf("%d: input %q", i, test.input)
				t.Fatalf("expected parse error but did not fail")
			}
			continue
		}

		if err != nil {
			t.Logf("%d: input %q", i, test.input)
			t.Fatalf("unexpected error, %v", err)
		}

		require.NotEqual(t, err, errUnexpected, "unexpected error occurred")

		require.Equal(t, test.expected, out, "%d: input %q", i, test.input)
	}
}
