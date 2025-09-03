package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pongNew"})
	})

	router.GET("/api/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"greeting": "Hello from Go + GinD1"})
	})

	router.GET("/env", func(c *gin.Context) {
		dbUrl := os.Getenv("DATABASE_URL")
		c.JSON(http.StatusOK, gin.H{"DATABASE_URL": dbUrl})
	})

	router.Run(":8056")
}
