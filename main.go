package main

import (
	"net/http"
	"time"

	routes "MapleService/control"
)

func main() {
	router := routes.NewRouter()

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:32001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}
