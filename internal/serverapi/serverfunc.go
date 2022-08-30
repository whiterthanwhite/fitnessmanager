package serverapi

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/whiterthanwhite/fitnessmanager/internal/db"
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

func InsertRecord(ctx context.Context, conn *db.Conn) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var requestBody []byte
		var err error

		requestBody, err = ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, err.Error()+" cannot read request body", http.StatusInternalServerError)
			return
		}
		var record fitnessdata.Record
		if err = json.Unmarshal(requestBody, &record); err != nil {
			http.Error(rw, err.Error()+" cannot unmarshal request body", http.StatusInternalServerError)
			return
		}
		ct, err := conn.InsertRecord(ctx, &record)
		if err != nil {
			http.Error(rw, err.Error()+" cannot insert record", http.StatusInternalServerError)
			return
		}
		log.Printf("Function \"InsertRecord\"; rows affected: %v; info: %v", ct.RowsAffected, ct.Info)
	}
}
