package controllers

import (
	"github.com/gin-gonic/gin"
)

// POST /set/task
// Set Task
func SetTask(c *gin.Context) {

}

// POST /fetch/task
// Fetch Task
func FetchTask(c *gin.Context) {
	// var task []models.Task
	// db := models.ConnectDB()
	// db.Find(&task)
	// c.JSON(http.StatusOK, gin.H{"data": task})
}
