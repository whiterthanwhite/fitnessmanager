package fitnessdata

import (
	"time"
)

type Record struct {
	Date        time.Time `json:"date"`
	Name        string    `json:"name"`
	Take        int       `json:"take"`
	Repetitions int       `json:"repetitions"`
	Description string    `json:"description"`
}
