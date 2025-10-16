package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	User_id int `json:"user_id"`
	jwt.RegisteredClaims
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {

		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			context.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
			return []byte("secret_key"), nil
		})

		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			context.Abort()
			return
		}

		context.Set("user_id", claims.User_id)
		context.Next()

	}
}
