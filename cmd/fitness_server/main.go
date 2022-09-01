package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	db "github.com/whiterthanwhite/fitnessmanager/internal/db"
	"github.com/whiterthanwhite/fitnessmanager/internal/serverapi"
)

func main() {
	fmt.Println(time.Now())
	ctx := context.Background()

	// connect to database
	dbAddr := "postgresql://localhost:5432/postgres" // test
	conn, err := db.Connect(ctx, dbAddr)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Connected to database")
	defer conn.Close(ctx)

	tableExists, err := conn.TableExist(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	var commTag *db.CommandTag
	if tableExists {
		// commTag, err = conn.DropTables(ctx)
	} else {
		commTag, err = conn.InitTables(ctx)
	}
	if err != nil {
		log.Fatal(err.Error())
	}
	if commTag != nil {
		log.Println(commTag)
	}

	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err.Error())
	}
	sm := http.NewServeMux()
	sm.HandleFunc("/training/entry_no", serverapi.GetTrainingRecordByEntryNo(ctx, conn))
	sm.HandleFunc("/training/date", serverapi.GetTrainingRecordByDate(ctx, conn))
	sm.HandleFunc("/training/insert", serverapi.InsertRecord(ctx, conn))

	s := &http.Server{
		Handler: sm,
	}

	log.Printf("Listening on %s ...\n", s.Addr)
	if err = s.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Finished")
}
