package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()

	dsn := "postgresql://ppm@localhost:5432/sipman"
	db, err := gorm.Open(postgres)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run()
}
