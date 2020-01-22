package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Show(c *gin.Context) {
	c.Param("name")
	c.JSON(200, gin.H{
		"name": "tom",
		"sex":  "mancx",
	})
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome",
	})

}
func Log4j(param gin.LogFormatterParams) string {

	// your custom format
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\"\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		//param.Request.UserAgent(),
		param.ErrorMessage,
	)
}
