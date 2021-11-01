package httpapi

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	
	"github.com/devhg/es/handler"
)

func NewHTTPRouter() *gin.Engine {
	r := gin.Default()
	
	gin.ForceConsoleColor()
	
	// http 接口注册
	setAPIRouter(r)
	
	return r
}

func setAPIRouter(r *gin.Engine) {
	// 接口注册
	api := r.Group("/api")
	
	userApi := api.Group("/user")
	{
		userApi.Handle(http.MethodPost, "/create", handler.Create)
		userApi.Handle(http.MethodPut, "/update", handler.Update)
		userApi.Handle(http.MethodDelete, "/delete", handler.Delete)
		userApi.Handle(http.MethodGet, "/info", handler.MGet)
		userApi.Handle(http.MethodPost, "/search", handler.Search)
		
	}
}
