package util

import (
	"fmt"
	"regexp"
	"time"
)

const (
	formatDate    = "2006-01-02"
	layoutTime    = "2006-01-02 15:04:05"
	layoutNotDash = "200601021504"
)

func IsNumeric(s string) bool {
	re := regexp.MustCompile(`^[0-9]+$`)
	return re.MatchString(s)
}

func ToDateTime(time_ string, withDash bool) time.Time {
	var t time.Time
	if withDash {
		t, _ = time.Parse(layoutTime, time_)
	} else {
		t, _ = time.Parse(layoutNotDash, time_)
	}

	return t
}

func CompareDate(fileTime time.Time) bool {
	now := time.Now()
	d30 := now.Add(-time.Hour * 30 * 24)
	stringNow := fmt.Sprintf("%v 23:59:59", now.Format(formatDate))
	d30StringNow := fmt.Sprintf("%v 00:00:00", d30.Format(formatDate))
	if ToDateTime(d30StringNow, true).Before(fileTime) && ToDateTime(stringNow, true).After(fileTime) {
		return true
	}
	return false
}
