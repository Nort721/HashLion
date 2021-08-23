package utils

import (
	"strconv"
	"time"
)

// returns the amount of time between two dates
// with the time unit attached
func GetTimeBetweenDates(finishDate time.Time, startDate time.Time) string {

	timeDiff := finishDate.Sub(startDate)

	diff := timeDiff.Hours()

	unit := "hours"

	// ToDo: find a better way to do this
	if diff < 1 {

		diff = timeDiff.Minutes()

		unit = "minutes"

		if diff < 1 {

			unit = "seconds"

			diff = timeDiff.Seconds()

			if diff < 1 {

				unit = "milliseconds"

				diff = float64(timeDiff.Milliseconds())

				if diff < 1 {

					unit = "microseconds"

					diff = float64(timeDiff.Microseconds())
				}

			}
		}
	}

	return strconv.FormatFloat(diff, 'f', 1, 64) + " " + unit
}
