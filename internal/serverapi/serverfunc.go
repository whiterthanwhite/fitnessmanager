package serverapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/whiterthanwhite/fitnessmanager/internal/fitnessdata"
)

func GetTrainingRecord() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		t := getTraining()
		responce, err := json.Marshal(&t)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		rw.Write(responce)
	}
}

// test
func getTraining() *fitnessdata.Record {
	return &fitnessdata.Record{
		Date:        time.Now(),
		Name:        "test",
		Take:        1,
		Repetitions: 10,
		Description: "some description",
	}
}
