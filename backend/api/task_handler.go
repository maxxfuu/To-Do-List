package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxxfuu/To-Do-List/backend/database"
	"github.com/maxxfuu/To-Do-List/backend/models"
)

// Before invoking any client request, we need to validate JWT token.

// POST: To add a task into a to-do list
func postTask(c *gin.Context) {
	var tasks models.Task

	if err := c.ShouldBindJSON(&tasks); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invald JSON data"})
		return
	}

	// Check Tasks Title
	if tasks.Title == "" {
		log.Print("Input Error", "Task title required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task title requried"})
		return
	}

	query := "INSERT INTO tasks (user_id, title, content, status) VALUES ($1, $2, $3, $4) RETURNING task_id"
	if err := database.DB.QueryRow(query, tasks.UserID, tasks.Title, tasks.Content, tasks.Status).Scan(&tasks); err != nil {
		log.Printf("Error Inserting Tasks%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to save task"})
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Task added sucessfully", "task_id": tasks.TaskID})
}

// PUT: To update a task in to-do list
func putTask(c *gin.Context) {
	var task models.Task
	var err error
	err = c.ShouldBindJSON(&task)

	if err != nil {

	}

}

// DELETE: To delete a task
func deleteTask(c *gin.Context) {

}
