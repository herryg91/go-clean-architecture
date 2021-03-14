package helpers

import (
	"math"
	"time"
)

func CountAge(birthday time.Time) int {
	today := time.Now()
	return int(math.Floor(today.Sub(birthday).Hours() / 24 / 365))
}
