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

func CompareDate(bottomTime_ int, fileTime time.Time) bool {
	now := time.Now()
	stringNow := fmt.Sprintf("%v 23:59:59", now.Format(formatDate))
	bottomTime := now.Add(-time.Hour * time.Duration(bottomTime_) * 24)
	bottomTimeString := fmt.Sprintf("%v 00:00:00", bottomTime.Format(formatDate))
	if ToDateTime(bottomTimeString, true).Before(fileTime) && ToDateTime(stringNow, true).After(fileTime) {
		return true
	}

	return false
}
