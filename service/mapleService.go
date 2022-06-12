package service

import (
	"net/http"
	// "strconv"

	"github.com/gin-gonic/gin"

	provider "MapleService/provider"
)

type ApiResponse struct {
	ResultCode    string
	ResultMessage interface{}
}

func Test(c *gin.Context) {
	secondParameter := c.Query("second")
	response := provider.GetTest(secondParameter)
	// firstDefaultParameter := c.DefaultQuery("first", "預設")
	c.Data(http.StatusOK, "application/json", response)
	// vars := mux.Vars(r)
	// queryId := vars["id"] //獲取url參數

	// response, code := provider.GetTest(queryId)

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(code)
	// w.Write(response)
}
