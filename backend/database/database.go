package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver"
)

var DB *sqlx.DB // Global representation of a Database

func InitDB() {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"postgres",
		"",
		"localhost",
		"5432",
		"postgres")

	var err error

	// pointer to instance of db, err = open connection("driver name", )
	DB, err = sqlx.Open("postgres", connString)
	if err != nil {
		log.Fatalf("Failed to connect to database, Connection Failure: %v", err)
		return
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Database Unreachable, Ping Failure: %v", err)
	}

	log.Println("Connected to database")
}
