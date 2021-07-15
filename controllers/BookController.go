package controllers

import (
	"fmt"
	"go-gin-api-rest/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var input models.CreateBookModel

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// YYYY-MM-DD
	if input.Birthdate != nil {
		const layout = "2006-01-02"

		dateValue, err := time.Parse(layout, *input.Birthdate)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(dateValue)
		formatDate := dateValue.Format(layout)
		fmt.Println(formatDate)
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func FindBook(c *gin.Context) {
	birthdate := "2021-14-06"
	isActive := true
	book := models.CreateBookModel{
		ID:        45,
		Title:     "Programming",
		Author:    "Miguel",
		Sex:       "male",
		Status:    5,
		Birthdate: &birthdate,
		IsActive:  &isActive,
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func GetBookById(c *gin.Context) {
	idPath := c.Param("id")

	id, err := strconv.Atoi(idPath)
	if err == nil {
		isActive := true
		book := models.CreateBookModel{
			ID:       uint(id),
			Title:    "Programming",
			Author:   "Maria",
			Sex:      "female",
			Status:   4,
			IsActive: &isActive,
		}

		c.JSON(http.StatusOK, gin.H{"data": book})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id is not a number"})
	}
}
