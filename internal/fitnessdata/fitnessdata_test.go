package fitnessdata

import (
	"encoding/json"
	"testing"
	"time"
)

func TestFitnessData(t *testing.T) {
	record := &Record{
		Date:        time.Now(),
		Name:        "test",
		Take:        0,
		Repetitions: 0,
		Description: "test",
	}
	m, err := json.Marshal(record)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(string(m))
}
