package domain

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaims struct {
	ID    string `json:"_id"`
	Email string `json:"email"`
	jwt.StandardClaims
}
