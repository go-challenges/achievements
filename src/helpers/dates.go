package helpers

import (
	"errors"
	"time"
)

func FindPeriod(period interface{}) (time.Time, error) {
	var duration time.Time
	switch period {
	case "day":
		duration = time.Now().AddDate(0, 0, 1)
	case "month":
		duration = time.Now().AddDate(0, 1, 0)
	case "year":
		duration = time.Now().AddDate(1, 0, 0)
	case "week":
		duration = time.Now().AddDate(0, 0, 7)
	default:
		return time.Time{}, errors.New("Undefined period!")
	}

	return duration, nil
}
