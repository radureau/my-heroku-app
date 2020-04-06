package server

import (
	"github.com/gin-gonic/gin"
	"github.com/radureau/my-heroku-app-pkg/helloyou"
)

// Serve _
func Serve() error {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	helloyou.AddHandlers(r, "v1/helloyou")

	return r.Run()
}
