package controllers

import (
	"go-gin-api-rest/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var input models.CreateBookModel

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func FindBook(c *gin.Context) {
	book := models.CreateBookModel{
		ID:     45,
		Title:  "Programming",
		Author: "Miguel",
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func GetBookById(c *gin.Context) {
	idPath := c.Param("id")

	id, err := strconv.Atoi(idPath)
	if err == nil {
		book := models.CreateBookModel{
			ID:     uint(id),
			Title:  "Programming",
			Author: "Miguel",
		}

		c.JSON(http.StatusOK, gin.H{"data": book})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id is not a number"})
	}
}
