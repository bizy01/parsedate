package parsedate

import (
	"time"
	"testing"
	"github.com/stretchr/testify/require"
)

type dateTest struct {
	in           string
	out          int64
	fail          bool
}

var year, month, day = time.Now().Date()

var tz, _ = time.LoadLocation("Asia/Shanghai")

var ts = time.Date(year, month, day, 0, 0, 0, 0, tz).Unix()

var testInputs = []dateTest{
	{in: "00:00:00", out: ts},
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
	{
		in: "aaa/04/02",
		out: 1617292800,
		fail: true,
	},
}

func TestParse(t *testing.T) {
	for i, th := range testInputs {
		expected, err := Timestamp(th.in)

		if th.fail {
			hasError := false

			if err != nil {
				hasError = true
			}

			if !hasError {
				t.Logf("%d: input %q", i, th.in)
				t.Fatalf("expected parse error but did not fail")
			}
			continue
		}

		if err != nil {
			t.Log(err)
		}

		require.Equal(t, th.out, expected, "%d: input %q", i, th.in)
	}
}