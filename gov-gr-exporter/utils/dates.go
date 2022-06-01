package utils

import (
	"log"
	"time"
)

// todo: tidy up to make adding more date formats easier

func DateStringToTimeZulu(date string) (time.Time) {
	timestamp, err := time.Parse("2006-01-02T15:04:05Z", date)
	if err != nil {
		log.Fatal(err)
	}
	return timestamp
}

func DateStringToTimePlainDate(date string) (time.Time) {
	timestamp, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Fatal(err)
	}
	return timestamp
}