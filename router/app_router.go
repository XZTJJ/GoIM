package router

import (
	"com.zhouhc/goim/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/index", service.GetIndex)
	r.POST("/user/createUser", service.UserCreate)
	return r
}
