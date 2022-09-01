package serverapi

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/whiterthanwhite/fitnessmanager/internal/db"
	"github.com/whiterthanwhite/fitnessmanager/internal/fitnessdata"
)

func GetTrainingRecordByEntryNo(ctx context.Context, conn *db.Conn) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		u := r.URL
		v := u.Query()
		if !v.Has("entry_no") {
			http.Error(rw, "value entry_no is missed", http.StatusBadRequest)
			return
		}
		entryNo, err := strconv.ParseInt(v.Get("entry_no"), 0, 64)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		record, err := conn.GetRecordByEntryNo(ctx, int(entryNo))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		responseBody, err := json.Marshal(&record)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		rw.Write(responseBody)
	}
}

func GetTrainingRecordByDate(ctx context.Context, conn *db.Conn) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		u := r.URL
		v := u.Query()
		if !v.Has("entry_no") {
			http.Error(rw, "value entry_no is missed", http.StatusBadRequest)
			return
		}
		date, err := time.Parse("", v.Get("date"))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		records, err := conn.GetRecordByDate(ctx, date)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		responseBody, err := json.Marshal(records)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		rw.Write(responseBody)
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
