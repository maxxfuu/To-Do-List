// Main Entry Point

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/maxxfuu/To-Do-List/backend/middleware"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	// auth_handler.go
	router.POST("/signin", postSignIn)
	router.POST("/signup", postSignUp)
	router.DELETE("/users/:username", deleteUser)

	// task_handler.go
	router.POST("/task", middleware.JWTAuthentication(), postTask)
	router.PUT("/task", middleware.JWTAuthentication(), putTask)
	router.DELETE("/task", middleware.JWTAuthentication(), deleteTask)

	return router
}
