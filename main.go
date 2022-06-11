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
		Handler:        router,
		Addr:           "0.0.0.0:23001",
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(srv.ListenAndServe())
}
