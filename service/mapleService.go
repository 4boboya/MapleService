package service

import (
	"net/http"

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
	c.Data(http.StatusOK, "application/json", response)
}
