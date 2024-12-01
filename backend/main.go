package main

import (
	"log"

	"github.com/maxxfuu/To-Do-List/backend/api"
	"github.com/maxxfuu/To-Do-List/backend/database"
)

func main() {

	database.InitDB()
	log.Println("Database Initialized")

	router := api.SetUpRouter()
	log.Println("HTTP server starting on PORT 8080...  ")
	router.Run(":8080")
}
