package tokenutil

import (
	"time"

	"github.com/ETjoel/task_managment_api/domain"
	"github.com/dgrijalva/jwt-go"
)

func CreateAccessToken(jwtsecret string, user domain.User, expiryHour int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiryHour))

	

	claims := &jwt.MapClaims{
		"_id": user.ID,
		"email": user.Email,
		jwt.StandardClaims{
			ExpirsAt: exp
		}
	}

}