package convbox

import (
	"strconv"
	"strings"
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
			return number[0:len(number)-6] + "ล้าน"
		}
		return m + "." + m1 + "ล้าน"
	}

	if len(number) > 3 {
		m := number[0 : len(number)-3]
		m1 := number[len(number)-3 : len(number)-2]
		if m1 == "0" {
			return m + "พัน"
		}
		return m + "." + m1 + "พัน"
	}

	return number

}

func ShortNumberWorld(count int) string {

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

func ReFormatTime(t int64) string {

	var sTime string
	now := time.Now().Unix()

	duration := now - t

	if duration < 60 {
		sDuration := strconv.FormatInt(duration, 10)
		sTime = sDuration + " วินาทีที่แล้ว"
	} else if duration < 3600 {
		mDuration := strconv.FormatInt(duration/60, 10)
		sTime = mDuration + " นาที"
	} else if duration < 86400 {
		hDuration := strconv.FormatInt(duration/3600, 10)
		sTime = hDuration + " ชั่วโมง"
	}
	if duration >= 86400 {
		formatTime := time.Unix(t, 0).Format("2006-01-02 15:04:05")
		values := strings.Split(formatTime, "-")
		if values[1] != "" {

			var mouth string
			switch values[1] {
			case "01":
				mouth = "มกราคม"
			case "02":
				mouth = "กุมภาพันธ์"
			case "03":
				mouth = "มีนาคม"
			case "04":
				mouth = "เมษายน"
			case "05":
				mouth = "พฤษภาคม"
			case "06":
				mouth = "มิถุนายน"
			case "07":
				mouth = "กรกฎาคม"
			case "08":
				mouth = "สิงหาคม"
			case "09":
				mouth = "กันยายน"
			case "10":
				mouth = "ตุลาคม"
			case "11":
				mouth = "พฤศจิกายน"
			case "12":
				mouth = "ธันวาคม"
			}

			dayTime := strings.Split(values[2], " ")
			sTime = dayTime[0] + " " + mouth + " " + values[0] + " " + dayTime[1]

		} else {
			sTime = formatTime
		}

	}

	return sTime
}

func FormatTime(t int64) string {

	var sTime string

	formatTime := time.Unix(t, 0).Format("2006-01-02 15:04:05")
	values := strings.Split(formatTime, "-")
	if values[1] != "" {

		var mouth string
		switch values[1] {
		case "01":
			mouth = "มกราคม"
		case "02":
			mouth = "กุมภาพันธ์"
		case "03":
			mouth = "มีนาคม"
		case "04":
			mouth = "เมษายน"
		case "05":
			mouth = "พฤษภาคม"
		case "06":
			mouth = "มิถุนายน"
		case "07":
			mouth = "กรกฎาคม"
		case "08":
			mouth = "สิงหาคม"
		case "09":
			mouth = "กันยายน"
		case "10":
			mouth = "ตุลาคม"
		case "11":
			mouth = "พฤศจิกายน"
		case "12":
			mouth = "ธันวาคม"
		}

		dayTime := strings.Split(values[2], " ")
		sTime = dayTime[0] + " " + mouth + " " + values[0]

	}

	return sTime
}
