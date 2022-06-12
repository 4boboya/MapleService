package control

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	service "MapleService/service"
)

// type Route struct {
// 	Method     string
// 	Pattern    string
// 	Handler    http.HandlerFunc
// 	Middleware mux.MiddlewareFunc
// }

type Controller struct {
	*gin.Engine
}

func NewController(e *gin.Engine) *Controller {
	return &Controller{e}
}

//設定Router
func (r *Controller) Router() {
	api := r.Group("/api")
	{
		api.GET("/version", version)
		v1 := api.Group("/v1")
		{
			v1.GET("/param/:first", Param)
			v1.GET("/query", service.Test)
		}
	}
}

// @Summary 版本查詢
// @Tags Version
// @version 1.0
// @Router /version [get]
func version(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("%v V:%v-%v,Build:%v", "demo", "1.0.1", "Local", "2022/01/01 15:32:10"))
}

// @Summary Param 範例
// @Tags Param
// @version 1.0
// @Router /param/test [get]
func Param(c *gin.Context) {
	firstParameter := c.Param("first")
	c.String(http.StatusOK, "Param = %v\n", firstParameter)
}

// @Summary Query 範例
// @Tags Query
// @version 1.0
// @Router /query [get]
// func Test(c *gin.Context) {
// 	secondParameter := c.Query("second")
// 	response := ApiResponse{"200", ResponseData{Id: secondParameter, Data: 123}}
// 	// firstDefaultParameter := c.DefaultQuery("first", "預設")
// 	responseData, _ := json.Marshal(response)
// 	c.Data(http.StatusOK, "application/json", responseData)
// }

// var routes []Route

// func init() {
// 	registry("GET", "/test/{id}", service.Test, nil)
// }

// func NewRouter() *mux.Router {
// 	r := mux.NewRouter()
// 	for _, route := range routes {
// 		r.Methods(route.Method).
// 			Path(route.Pattern).
// 			Handler(route.Handler)
// 		if route.Middleware != nil {
// 			r.Use(route.Middleware)
// 		}
// 	}
// 	return r
// }

// func registry(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
// 	routes = append(routes, Route{method, pattern, handler, middleware})
// }
