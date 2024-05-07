package dateka

import (
	"math/rand"
	"time"
)

func RandomDate() time.Time {
	year := rand.Intn(2024)
	month := time.Month(rand.Intn(13))
	day := rand.Intn(32)
	hour := rand.Intn(25)
	minSec := func() int {
		return rand.Intn(60)
	}

	return time.Date(year, month, day, hour, minSec(), minSec(), 0, time.UTC)
}
