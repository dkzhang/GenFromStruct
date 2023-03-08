package utils

import (
	"strconv"
	"time"
)

func ConvertDate(dateInt int) (days int, err error) {
	date, err := time.Parse("20060102", strconv.Itoa(dateInt))
	if err != nil {
		return 0, err
	}

	return int(date.Unix() / 86400), nil
}
