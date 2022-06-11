package main

import (
	"fmt"
	"net/http"

	routes "MapleService/control"
)

func main() {
	fmt.Println("start")
	router := routes.NewRouter()

	http.ListenAndServe("0.0.0.0:23001", router)
}
