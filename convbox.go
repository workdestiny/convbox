package convbox

import (
	"strconv"
	"time"
)

const (
	timeFormat = "2006-01-02"
)

//ShortNumber use convert number (integer) to short number (string)
func ShortNumber(count int) string {

	number := strconv.Itoa(count)
	if len(number) > 6 {
		m := number[0 : len(number)-6]
		m1 := number[len(number)-6 : len(number)-5]
		if m1 == "0" {
			return number[0:len(number)-6] + "M"
		}
		return m + "." + m1 + "M"
	}

	if len(number) > 3 {
		m := number[0 : len(number)-3]
		m1 := number[len(number)-3 : len(number)-2]
		if m1 == "0" {
			return m + "K"
		}
		return m + "." + m1 + "K"
	}

	return number

}

// GetYMD return 3 Value (Year, Month, Day)
func GetDate(t time.Time) (string, string, string) {
	date := t.Format(timeFormat)
	return date[0:4], date[5:7], date[8:10]
}

// GetYear return Year
func GetYear(t time.Time) string {
	date := t.Format(timeFormat)
	return date[0:4]
}

// GetMonth return Month
func GetMonth(t time.Time) string {
	date := t.Format(timeFormat)
	return date[5:7]
}

// GetDay return Day
func GetDay(t time.Time) string {
	date := t.Format(timeFormat)
	return date[8:10]
}
