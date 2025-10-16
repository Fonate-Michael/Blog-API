package controller

import (
	"app/db"
	"app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

func Register(context *gin.Context) {
	var req models.User

	err := context.BindJSON(&req)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind json ma boi"})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate password hash"})
		return
	}

	_, err = db.DB.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", req.Username, req.Email, hashPassword)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})

}

func Login(context *gin.Context) {
	var req models.User

	err := context.BindJSON(&req)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Failed to bind json ma boi"})
		return
	}

	var user models.User
	err = db.DB.QueryRow("SELECT * FROM users WHERE email = $1", req.Email).Scan(&user.Id, &user.Username, &user.Email, &user.Password)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to query user"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to compare password"})
		return
	}

	expiration := time.Now().Add(1 * time.Hour)

	claims := &Claims{
		UserId: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var jwtKey = []byte("secret_key")

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"token": tokenString})

}
