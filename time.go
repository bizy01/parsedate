package parsedate

import (
	"time"
	"parsedate/parser"
)

// to time
func Parse(datetime string) (time.Time, error) {
	node, err := parser.ParseDate(datetime)
	if err != nil {
		return time.Time{}, err
	}

	dt := node.(*parser.Date)
	tm := time.Date(dt.Year, time.Month(dt.Month), dt.Day, 0, 0, 0, 0, time.UTC)

	return tm, nil
}

// to timestamp
func Timestamp(datetime string) (int, error) {
	tm, err := Parse(datetime)
	if err != nil {
		return -1, err
	}

	return tm.Unix(), nil
}