package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salty-ivy/chatsa/database"
)

func main() {

	// initiate DataBase
	database.ConnectSQLite()

	// initiate gin router
	r := gin.Default()

	// Simple ping endpoint for 3P services for server checks.
	// use Anon func as instant handler
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.Run("localhost:8000")
}
