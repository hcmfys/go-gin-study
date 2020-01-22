package main

import (
	"github.com/gin-gonic/gin"
	"router"
)

func run() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	gin.ForceConsoleColor()
	router.AddRouter(r)
	r.Run(":1983")
}

func main() {
	run()
}
