package provider

import (
	"encoding/json"
)

type ApiResponse struct {
	ResultCode    string
	ResultMessage interface{}
}

type ResponseData struct {
	Id   string
	Data int
}

func GetTest(id string) ([]byte, int) {
	response := ApiResponse{"200", ResponseData{Id: id, Data: 123}}

	responseData, _ := json.Marshal(response)

	return responseData, 200
}
