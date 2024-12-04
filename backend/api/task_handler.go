package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxxfuu/To-Do-List/backend/models"
)

// Before invoking any client request, we need to validate JWT token.

// POST: To add a task into a to-do list
func postTask(c *gin.Context) {
	var tasks models.Task
	err := c.ShouldBindJSON(&tasks)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid JSON Data": err})
	}
}

// PUT: To update a task in to-do list
func putTask(c *gin.Context) {

}

// DELETE: To delete a task
func deleteTask(c *gin.Context) {

}
