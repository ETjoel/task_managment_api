package tokenutil

import (
	"errors"
	"fmt"
	"time"

	"github.com/ETjoel/task_managment_api/domain"
	"github.com/dgrijalva/jwt-go"
)

func CreateAccessToken(jwtsecret string, user domain.User, expiryHour int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiryHour))

	claims := &domain.JwtCustomClaims{
		ID:    user.ID,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err := token.SignedString([]byte(jwtsecret))
	if err != nil {
		return "", errors.New("internal server error: " + err.Error())
	}

	return jwtToken, nil
}

func ExtractEmail(requestToken string, jwtSecret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	return claims["email"].(string), nil
}
