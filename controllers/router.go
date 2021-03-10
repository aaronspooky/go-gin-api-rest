package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRequests() *gin.Engine {
	log.Println("Start server at localhost")
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello World!"})
	})

	r.POST("/createbook", CreateBook)
	r.GET("/getBook", FindBook)
	r.GET("/getBook/:id", GetBookById)

	return r
}
