package server

import "github.com/gin-gonic/gin"

// Serve _
func Serve() error {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
	return r.Run()
}