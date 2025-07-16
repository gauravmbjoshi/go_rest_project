package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)
const secretKey = "super-secret"
func GenerateToken(email string,userId int64) (string,error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":email,
		"userId":userId,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
		// exp is at what time it should expire so we added 2 hours from the time of login and we get value in Unix
	})
	// this will create a token with claims meaning data attached to it
	return token.SignedString([]byte(secretKey))
}

func CheckToken(token string) (int64,error) {
	parsedToken,err := jwt.Parse(token,func (token *jwt.Token)(interface{},error){
		_,ok := token.Method.(*jwt.SigningMethodHMAC) // to check if the token has same SigningMethodHS256 which we used to generate token
		// in go if you want to check the type you can add . and then in parenthesis the type you want it to be so in above line Method.(*jwt.SigningMethodHMAC) we check the type of the token.Method
		if !ok {
			return nil,errors.New("Unexpected signing method")
		}
		return []byte(secretKey),nil
	})
	if err != nil {
		return 0,errors.New("could not parse token")
	}
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid{
		return 0,errors.New("Invalid token")
	}
	claims ,ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0,errors.New("Invalid token claim")
	}
	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64)) // here we checked if it is float64 and then converted to int64
	// the above code is to showcase how you can get the email and id which you created in jwt while login
	return userId,nil

}