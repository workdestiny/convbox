package convbox

import "strconv"

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
