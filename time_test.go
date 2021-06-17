package parsedate

import (
	"testing"
	"github.com/stretchr/testify/require"
)

type dateTest struct {
	in           string
	out          int64
	err          bool
}

var testInputs = []dateTest{
	{in: "10:20:30", out: 1623896430},
	{in: "2021-1-1", out: 1609430400},
	{in: "2021-1-1 10:20:30", out: 1609467630},
	{in: "2021-01-01 10:20:30", out: 1609467630},
	{in: "oct 7, 1970 10:20:30", out: 24114030},
	{in: "oct. 7, 1970 10:20:30", out: 24114030},
	{in: "sept. 7, 1970 10:20:30", out: 21522030},
	{in: "sept. 7, 1970 10:20:30", out: 21522030},
	{in: "sept. 7, 1970 10:20:30", out: 21522030},
	{in: "7 oct 70", out: 24076800},
	{in: "7 oct 1970", out: 24076800},
	{in: "7 oct 1970", out: 24076800},
	{in: "03 February 2013", out: 1359820800},
	{in: "2021年10月1日", out: 1633017600},
	{in: "03/31/2014", out: 1396195200},
	{in: "2021/04/02", out: 1617292800},
}

func TestParse(t *testing.T) {
	for i, th := range testInputs {
		expected, err := Timestamp(th.in)
		if err != nil {
			t.Log(err)
			continue
		}

		require.Equal(t, th.out, expected, "%d: input %q", i, th.in)
	}
}