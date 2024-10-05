package middleware

import (
	"fmt"
	"strings"

	"github.com/ETjoel/task_managment_api/data"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.IndentedJSON(401, gin.H{"error": "authorization needed"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.IndentedJSON(401, gin.H{"error": "invalid authorization"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return data.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.IndentedJSON(401, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

	}
}
