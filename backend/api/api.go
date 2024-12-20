// Main Entry Point

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/maxxfuu/To-Do-List/backend/middleware"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	// auth_handler.go
	router.POST("/signup", postSignUp)
	router.POST("/signin", postSignIn)
	router.DELETE("/users/:username", middleware.JWTAuthentication(), deleteUser) // protected routes

	// task_handler.go
	router.POST("/task", middleware.JWTAuthentication(), postTask)     // protected routes
	router.PUT("/task", middleware.JWTAuthentication(), putTask)       // protected routes
	router.DELETE("/task", middleware.JWTAuthentication(), deleteTask) // protected routes

	return router
}
