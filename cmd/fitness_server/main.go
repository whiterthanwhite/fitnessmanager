package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/whiterthanwhite/fitnessmanager/internal/serverapi"
)

func main() {
	fmt.Println(time.Now())
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err.Error())
	}
	sm := http.NewServeMux()
	sm.HandleFunc("/", serverapi.GetTrainingRecord())

	s := &http.Server{
		Handler: sm,
	}
	if err = s.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
