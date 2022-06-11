package main

import (
	"net/http"

	routes "MapleService/control"
)

func main() {
	router := routes.NewRouter()
	http.ListenAndServe(":32001", router)
}
