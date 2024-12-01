package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	// "github.com/jmoiron/sqlx"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/maxxfuu/To-Do-List/backend/database"
	"github.com/maxxfuu/To-Do-List/backend/models"
)

var secretKey = []byte("secret-key")

func createToken(username string, password string) (string, error) {
	//		create a new JWT token.(signing method, relevant info)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"password": password,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	//					sign the token with secret key and return generate token as a string
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if token.Valid != true {
		return fmt.Errorf("invalid token")

	}

	return nil
}

func postSignIn(c *gin.Context) {
	var user models.Users
	err := c.ShouldBindJSON(&user)

	// Check If Data Binding Is Possible
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		// To Log Errors
		log.Println("error", err)
		return
	}

	// Check if input value from client exists
	if user.Username == "" && user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required inputs"})
		return
	}

	// Make a query
	query := "SELECT username, password FROM users WHERE username = $1 AND password = $2 LIMIT 1"
	// Execute Query
	row := database.DB.QueryRow(query, user.Username, user.Password)
	var dbUsername, dbPassword string
	err = row.Scan(dbUsername, dbPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		} else {
			log.Printf("Database Error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"Internal Database Error:": err})

		}
	}

	// If query is successful, generate JWT
	token, err := createToken(user.Username, user.Password)
	if err == nil {
		c.JSON(http.StatusCreated, token)
	}

}

func postSignUp(c *gin.Context) {

}
