package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	router.POST("/log-in", LogInHandler)
	router.GET("/verify", VerifyHandler)

	router.Run(":8080")
}

type LogInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LogInHandler(c *gin.Context) {
	var req *LogInRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid request")
		return
	}

	if req.Username != "semyeong" || req.Password != "000130" {
		c.String(http.StatusUnauthorized, "Incorrect username or password")
		return
	}

	token, err := CreateToken(req.Username)
	if err != nil {
		c.String(http.StatusInternalServerError, "CreateToken error")
		return
	}

	c.String(http.StatusOK, token)
}

func VerifyHandler(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.String(http.StatusUnauthorized, "Missing authorization header")
		return
	}

	tokenString = tokenString[len("Bearer "):]

	err := VerifyToken(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, "Invalid Token")
		return
	}

	c.String(http.StatusOK, "Valid Token!")
}

var secretKey = []byte("secret-key")

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
