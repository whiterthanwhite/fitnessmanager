package serverapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/whiterthanwhite/fitnessmanager/internal/db"
)

func TestGetTrainingRecordByEntryNo(t *testing.T) {

}

func TestGetTrainingRecordsByDate(t *testing.T) {

}

func TestInsertRecord(t *testing.T) {
	ctx := context.Background()
	conn, err := db.Connect(ctx, "postgresql://localhost:5432/postgres")
	require.Nil(t, err)
	require.NotNil(t, conn)
	defer conn.Close(ctx)

	requestBody := bytes.NewBuffer([]byte(`{"date":"2022-08-31T01:09:38.403497+03:00","name":"test","take":0,"repetitions":0,"description":"test"}`))
	request := httptest.NewRequest(http.MethodPost, "localhost:8080/insert", requestBody)
	w := httptest.NewRecorder()
	h := http.HandlerFunc(InsertRecord(ctx, conn))
	h.ServeHTTP(w, request)
	result := w.Result()

	resultBody, err := ioutil.ReadAll(result.Body)
	require.Nil(t, err)
	t.Log(result.StatusCode)
	t.Log(string(resultBody))
}
