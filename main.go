package main

import (
	"go-gin-api-rest/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello World!"})
	})

	r.POST("/createbook", controllers.CreateBook)
	r.GET("/getBook", controllers.FindBook)
	r.GET("/getBook/:id", controllers.GetBookById)

	r.Run()
}
