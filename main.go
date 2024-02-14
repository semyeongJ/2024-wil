package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	router.GET("/secret", basicAuthMiddleware, func(c *gin.Context) {
		c.String(200, "success!")
	})

	router.Run(":8080")
}

func basicAuthMiddleware(c *gin.Context) {
	username, password, ok := c.Request.BasicAuth()

	if !ok || username != "semyeong" || password != "000130" {
		c.AbortWithStatus(401)
		return
	}

	c.Next()
}
