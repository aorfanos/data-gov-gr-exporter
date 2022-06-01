package utils

import (
	"log"
	"time"
)

func DateStringToTime(date string) (time.Time) {
	timestamp, err := time.Parse("2006-01-02T15:04:05Z", date)
	if err != nil {
		log.Fatal(err)
	}
	return timestamp
}