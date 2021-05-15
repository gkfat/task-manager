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
	c.BindJSON(&reqBody)

	db := data.ConnectDB()
	db.Create(&dbTable.Task{
		Title:       reqBody.Title,
		Description: reqBody.Description,
		Category:    reqBody.Category,
		Month:       reqBody.Month,
		Week:        reqBody.Week,
		Weekday:     reqBody.Weekday,
		WorkingHour: reqBody.WorkingHour,
	})
	var task api.Task
	db.Last(&task)

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
	c.BindJSON(&reqBody)

	var tasks []api.Task
	var result []api.Task
	var count int64

	db := data.ConnectDB()
	query := db.Find(&tasks)
	query.Count(&count)
	db.Limit(reqBody.Per).Offset(reqBody.Page).Find(&result)

	c.JSON(http.StatusOK, gin.H{
		"Code":    http.StatusOK,
		"Success": true,
		"Data":    tasks,
		"Count":   count,
		"Page":    reqBody.Page,
		"Per":     reqBody.Per,
	})
}
