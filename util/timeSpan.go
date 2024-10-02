package util

import (
	"errors"
	"time"
)

func ParseTimeSpan(startString string, endString string) (time.Time, time.Time, error) {
	var start, end time.Time
	var errStart, errEnd error
	if startString == "" {
		start = time.Date(1970, 1, 1, 0, 0, 1, 0, &time.Location{})
	} else {
		start, errStart = time.Parse("2006-01-02", startString)
	}
	if endString == "" {
		end = time.Now().UTC()
	} else {
		end, errEnd = time.Parse("2006-01-02", endString)
	}
	if errStart != nil || errEnd != nil {
		return start, end, errors.New("invalid start or end time format")
	}
	return start, end, nil
}
