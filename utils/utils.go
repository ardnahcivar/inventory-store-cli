package utils

import (
	"fmt"
	"time"
)

func ErrorHandler(e error) {

	if e != nil {
		fmt.Println(e)
	}
}

func ConvertStringToTime(timeString string) time.Time {
	timelLayout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(timelLayout, timeString)

	if err != nil {
		ErrorHandler(err)
	}

	return parsedTime

}
