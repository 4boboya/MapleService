package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	routes "MapleService/control"
)

func main() {
	fmt.Println("start")
	router := routes.NewRouter()

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:23001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
