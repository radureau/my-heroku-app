package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/radureau/my-heroku-app-pkg/helloyou"
)

var counter = 0 // subject to concurrent access

func incMW(c *gin.Context) {
	counter++
	c.Header("nth-request", fmt.Sprintf("%d", counter))
	c.Next()
}

// Serve _
func Serve() error {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	helloyou.AddHandlers(r, "helloyou/v1").Use(incMW)

	return r.Run()
}
