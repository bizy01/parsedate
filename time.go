package parsedate

import (
	"time"
	"parsedate/parser"
)

// to time
func Parse(datetime string) (time.Time, error) {
	dt, err := parser.ParseDate(datetime)
	if err != nil {
		return time.Time{}, err
	}

	year := dt.Year
	month := time.Month(dt.Month)
	day := dt.Day

	if year == 0 && month ==0 && day == 0 {
		year, month, day = time.Now().Date()
	}

	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	tm := time.Date(year, month, day, dt.Hour, dt.Minute, dt.Second, dt.Millisecond, beijing)

	return tm, nil
}

// to timestamp
func Timestamp(datetime string) (int64, error) {
	tm, err := Parse(datetime)
	if err != nil {
		return -1, err
	}

	return tm.Unix(), nil
}