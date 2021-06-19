package parsedate

import (
	"time"
	"github.com/bizy01/parsedate/parser"
)

var timezoneList = map[string]string{
	"-11":    "Pacific/Midway",
	"-10":    "Pacific/Honolulu",
	"-9:30":  "Pacific/Marquesas",
	"-9":     "America/Anchorage",
	"-8":     "America/Los_Angeles",
	"-7":     "America/Phoenix",
	"-6":     "America/Chicago",
	"-5":     "America/New_York",
	"-4":     "America/Santiago",
	"-3:30":  "America/St_Johns",
	"-3":     "America/Sao_Paulo",
	"-2":     "America/Noronha",
	"-1":     "America/Scoresbysund",
	"+0":     "Europe/London",
	"+1":     "Europe/Vatican",
	"+2":     "Europe/Kiev",
	"+3":     "Europe/Moscow",
	"+3:30":  "Asia/Tehran",
	"+4":     "Asia/Dubai",
	"+4:30":  "Asia/Kabul",
	"+5":     "Asia/Samarkand",
	"+5:30":  "Asia/Kolkata",
	"+5:45":  "Asia/Kathmandu",
	"+6":     "Asia/Almaty",
	"+6:30":  "Asia/Yangon",
	"+7":     "Asia/Jakarta",
	"+8":     "Asia/Shanghai",
	"+8:45":  "Australia/Eucla",
	"+9":     "Asia/Tokyo",
	"+9:30":  "Australia/Darwin",
	"+10":    "Australia/Sydney",
	"+10:30": "Australia/Lord_Howe",
	"+11":    "Pacific/Guadalcanal",
	"+12":    "Pacific/Auckland",
	"+12:45": "Pacific/Chatham",
	"+13":    "Pacific/Apia",
	"+14":    "Pacific/Kiritimati",
}

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

	tz, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return time.Time{}, err
	}

	tm := time.Date(year, month, day, dt.Hour, dt.Minute, dt.Second, dt.Millisecond, tz)

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