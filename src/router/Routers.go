package router

import "github.com/gin-gonic/gin"

func AddRouter(r *gin.Engine) {
	//r.Use(gin.LoggerWithFormatter(Log4j))
	r.Use(gin.Recovery())
	r.GET("/", Index)
	r.GET("/ping", Ping)
	r.GET("/show", Show)
}
