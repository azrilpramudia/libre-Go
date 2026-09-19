package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status":"ok"})
	})

	log.Println("Server Running on :5000")
	if err := r.Run(":5000"); err != nil {
		log.Fatal(err)
	}
}