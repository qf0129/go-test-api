package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "ok")
	})
	r.GET("/health", func(c *gin.Context) {
		log.Printf("health")
		c.JSON(200, gin.H{"status": "ok"})
	})
	r.GET("/api/hello", func(c *gin.Context) {
		log.Printf("hello1")
		println("hello2")
		c.JSON(200, "hello")
	})
	r.Run("0.0.0.0:8001")
}
