package dates

import "time"

const ParseDatePattern string = "2006-01-02"

func ParseDate(d string) (time.Time, error) {
	return time.Parse(ParseDatePattern, d)
}