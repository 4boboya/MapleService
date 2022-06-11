package service

import (
	"net/http"
	// "strconv"

	"github.com/gorilla/mux"

	provider "MapleService/provider"
)

type Todo struct {
	Id   int64
	Item string
}

type ApiResponse struct {
	ResultCode    string
	ResultMessage interface{}
}

var TodoList []Todo

func Test(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryId := vars["id"] //獲取url參數

	response, code := provider.GetTest(queryId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
