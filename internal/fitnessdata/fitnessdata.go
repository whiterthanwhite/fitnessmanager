package fitnessdata

import (
	"time"
)

type Record struct {
	Date        time.Time
	Name        string
	Take        int
	Repetitions int
	Description string
}
