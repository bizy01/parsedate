package parser

import (
	"strings"
	"github.com/spf13/cast"
)

type PositionRange struct {
	Start Pos
	End   Pos
}

type Node interface {
	Print() string
	PositionRange() *PositionRange
}

type DateTime struct {
    Date *Date
    Time *Time
}

func (d *DateTime) Print() string {
	return ""
}

func (d *DateTime) PositionRange() *PositionRange {
	return nil
}

type Date struct {
    Year int
    Month int
    Day int
}

func (d *Date) Print() string {
	return ""
}

func (d *Date) PositionRange() *PositionRange {
	return nil
}

type Time struct {
    Hour int
    Minute int
    Second int
    Millisecond int
}

func (t *Time) Print() string {
	return ""
}

func (t *Time) PositionRange() *PositionRange {
	return nil
}

type Year int

func JoinYear(tokens ...Token) Year {
	var out string
    for _, token := range tokens {
        out += token.Val
    }

    val := cast.ToInt(out)
    if val >= 70 && val < 100 {
    	val = 1900 + val
    } else if val < 70 {
        val = 2000 + val
    }

	return Year(val)
}

func (y Year) Print() string {
	return ""
}

func (y Year) PositionRange() *PositionRange{
	return nil
}

type Month int

func JoinMonth(tokens ...Token) Month {
	var out string
    for _, token := range tokens {
        out += token.Val
    }

	return Month(cast.ToInt(out))
}

func (m Month) Print() string {
	return ""
}

func (m Month) PositionRange() *PositionRange{
	return nil
}

type Day int

func JoinDay(tokens ...Token) Day {
	var out string
    for _, token := range tokens {
        out += token.Val
    }

	return Day(cast.ToInt(out))
}

func (d Day) Print() string {
	return ""
}

func (d Day) PositionRange() *PositionRange{
	return nil
}

type Hour int

func JoinHour(tokens ...Token) Hour {
	var out string
    for _, token := range tokens {
        out += token.Val
    }

    return Hour(cast.ToInt(out))
}

func (h Hour) Print() string {
	return ""
}

func (h Hour) PositionRange() *PositionRange{
	return nil
}

type Minute int

func JoinMinute(tokens ...Token) Minute {
	var out string
    for _, token := range tokens {
        out += token.Val
    }

    return Minute(cast.ToInt(out))
}

func (m Minute) Print() string {
	return ""
}

func (m Minute) PositionRange() *PositionRange{
	return nil
}

type Second int

func JoinSecond(tokens ...Token) Second {
	var out string
    for _, token := range tokens {
        out += token.Val
    }

    return Second(cast.ToInt(out))
}

func (s Second) Print() string {
	return ""
}

func (s Second) PositionRange() *PositionRange{
	return nil
}

var monthWords = map[string]int {
    "january":      1,
    "february":     2,
    "march":        3,
    "april":        4,
    "may":          5,
    "june":         6,
    "july":         7,
    "august":       8,
    "september":    9,
    "october":      10,
    "november":     11,
    "december":     12,
    "jan":          1,
    "feb":          2,
    "mar":          3,
    "apr":          4,
    "jun":          6,
    "jul":          7,
    "aug":          8,
    "sept":         9,
    "oct":          10,
    "nov":          11,
    "dec":          12,
}

func ConvToManth(m string) Month {
	if month, ok := monthWords[strings.ToLower(m)]; ok {
		return Month(month)
	}

	return Month(0)
}