package controllers

import (
	"go/data"
	"go/models/api"
	"go/models/dbTable"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /set/task
// Set Task
func SetTask(c *gin.Context) {
	var reqBody api.Task

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
	err = db.Create(&dbTable.Task{
		Title:       reqBody.Title,
		Description: reqBody.Description,
		CategoryID:  reqBody.CategoryID,
		LimitDate:   reqBody.LimitDate,
	}).Error
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"Code":    http.StatusBadRequest,
			"Success": false,
			"Data":    "Database error.",
		})
		return
	}
	var task api.Task
	db.Last(&task)
	db.Commit()

	c.JSON(http.StatusOK, gin.H{
		"Code":    http.StatusOK,
		"Success": true,
		"Data":    task.ID,
	})
}

// POST /fetch/task
// Fetch Task
func FetchTask(c *gin.Context) {
	var reqBody api.TaskQuery
	err := c.BindJSON(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Code":    http.StatusBadRequest,
			"Success": false,
			"Data":    "Incorrect request body.",
		})
		return
	}

	var tasks []api.Task
	var result []api.Task
	var count int64

	db := data.ConnectDB()
	query := db.Find(&tasks)
	query.Count(&count)
	err = db.Order("id DESC").
		Offset(reqBody.Page * reqBody.Per).
		Limit(reqBody.Per).
		Find(&result).Error
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"Code":    http.StatusBadRequest,
			"Success": false,
			"Data":    "Database error.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Code":    http.StatusOK,
		"Success": true,
		"Data":    result,
		"Count":   count,
		"Page":    reqBody.Page,
		"Per":     reqBody.Per,
	})
}
