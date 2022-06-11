package main

import (
	"fmt"
	"net/http"

	routes "MapleService/control"
)

func main() {
	fmt.Println("start")
	router := routes.NewRouter()

	http.ListenAndServe(":23001", router)

	fmt.Println("end")
}
