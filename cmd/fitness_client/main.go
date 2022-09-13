package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println(time.Now())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	httpClient := &http.Client{}
	reqBody := bytes.NewReader([]byte(""))
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodPut, "http://localhost:8080/training/insert", reqBody)
	if err != nil {
		log.Fatal(err.Error())
	}
	resp, err := httpClient.Do(httpRequest)
	if err != nil {
		log.Fatal(err.Error())
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	log.Println(resp.Status)
	log.Println(resp.StatusCode)
	log.Println(string(respBody))

	<-ctx.Done()
}
