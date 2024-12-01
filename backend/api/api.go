// Main Entry Point

package api

import (
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	// auth_handler
	router.POST("signin", postSignIn) // *
	router.POST("/signup", postSignUp)

	// task_handler
	router.GET("/task", getTask)
	router.POST("/task", postTask)
	router.PUT("/task", putTask)
	router.DELETE("/task", deleteTask)

	// middleware

	return router
}
