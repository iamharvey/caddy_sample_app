package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Get greeting message.
	r.GET("/greeting", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "Good day from back-end server!"})
	})

	r.Run(":8080")
}


