package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/whiterthanwhite/fitnessmanager/internal/db"
)

func main() {
	fmt.Println(time.Now())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	dbAddr := "postgresql://localhost:5432/postgres" // test
	conn, err := db.Connect(ctx, dbAddr)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close(ctx)

	httpClient := &http.Client{}

	// Check last entry_no
	clientLastEntryNo, err := conn.GetLastEntryNo(ctx)
	// Check last entry_no on server
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodPut, "http://localhost:8080/training/lastEntryNo", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	respBody, err := doRequest(httpClient, httpRequest)
	if err != nil {
		log.Fatal(err.Error())
	}
	serverLastEntryNo, err := strconv.ParseInt(string(respBody), 10, 64)
	if err != nil {
		log.Fatal(err.Error())
	}
	if clientLastEntryNo == int(serverLastEntryNo) {
		log.Println("clientLastEntryNo = serverLastEntryNo")
	}
	if clientLastEntryNo != int(serverLastEntryNo) {
		log.Print("Last entry no ")
		log.Printf("Client %d; ", clientLastEntryNo)
		log.Printf("Server %d\n", serverLastEntryNo)
		if clientLastEntryNo > int(serverLastEntryNo) {
			log.Println("clientLastEntryNo > serverLastEntryNo")
		}
		if clientLastEntryNo < int(serverLastEntryNo) {
			log.Println("clientLastEntryNo < serverLastEntryNo")
		}
	}

	reqBody := bytes.NewReader([]byte(""))
	httpRequest, err = http.NewRequestWithContext(ctx, http.MethodPut, "http://localhost:8080/training/insert", reqBody)
	if err != nil {
		log.Fatal(err.Error())
	}
	respBody, err = doRequest(httpClient, httpRequest)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(string(respBody))

	<-ctx.Done()
}

func doRequest(client *http.Client, req *http.Request) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
