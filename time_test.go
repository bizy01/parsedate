package parsedate

import (
	"testing"
)

func TestParse(t *testing.T) {
	tm, err := Parse("2010-1-1")
	if err != nil {
		t.Log(err)
	}
	t.Log("tm ===>", tm)
}