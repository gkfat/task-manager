package controllers

import (
	"go/data"
	"go/models/api"
	"go/models/dbTable"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /set/category
// Set Category
func SetCategory(c *gin.Context) {
	var reqBody api.Category

	err := c.ShouldBindJSON(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Code":    http.StatusBadRequest,
			"Success": false,
			"Data":    "Incorrect request body.",
		})
		return
	}

	db := data.ConnectDB()
	db.Create(&dbTable.Category{
		Title: reqBody.Title,
	})

	var category api.Category
	db.Last(&category)
	db.Commit()

	c.JSON(http.StatusOK, gin.H{
		"Code":    http.StatusOK,
		"Success": true,
		"Data":    category.ID,
	})
}

// POST /fetch/category
// Fetch Category
func FetchCategory(c *gin.Context) {
	var categories []api.Category
	var count int64

	db := data.ConnectDB()
	query := db.Order("id DESC").Find(&categories)
	query.Count(&count)
	db.Commit()

	c.JSON(http.StatusOK, gin.H{
		"Code":    http.StatusOK,
		"Success": true,
		"Data":    categories,
		"Count":   count,
	})
}
