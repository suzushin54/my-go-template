package main

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping! pong!!",
		})
	})
	err := r.Run(":8083")
	if err != nil {
		log.Error(err)
	}
}
