package main

import (
	"fmt"
	"net/http"
	"time"

	routes "MapleService/control"
)

func main() {
	fmt.Println("start")
	router := routes.NewRouter()

	service := &http.Server{
		Handler:        router,
		Addr:           "0.0.0.0:23001",
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	service.ListenAndServe()
	fmt.Println("end")
}
