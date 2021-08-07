package router

import (
	"github.com/gin-gonic/gin"
	"shuwen/controller"
)

func Run()  {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.GET("/api/user/register", controller.Register)
	r.Run()
}
