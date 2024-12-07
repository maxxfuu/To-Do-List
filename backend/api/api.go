// Main Entry Point

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/maxxfuu/To-Do-List/backend/middleware"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	// auth_handler.go, protected routes
	router.POST("/signin", middleware.JWTAuthentication(), postSignIn)
	router.POST("/signup", middleware.JWTAuthentication(), postSignUp)
	router.DELETE("/users/:username", middleware.JWTAuthentication(), deleteUser)

	// task_handler.go
	router.POST("/task", middleware.JWTAuthentication(), postTask)
	router.PUT("/task", middleware.JWTAuthentication(), putTask)
	router.DELETE("/task", middleware.JWTAuthentication(), deleteTask)

	return router
}
