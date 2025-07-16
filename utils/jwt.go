package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)
const secretKey = "super-secret"
func GenerateToken(email string,id int64) (string,error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":email,
		"id":id,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
		// exp is at what time it should expire so we added 2 hours from the time of login and we get value in Unix
	})
	// this will create a token with claims meaning data attached to it
	return token.SignedString([]byte(secretKey))
}