package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string,error) {
	bytesPassword,err := bcrypt.GenerateFromPassword([]byte(password),14)
	return string(bytesPassword),err
}
func CheckPassword(password,hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
	return err == nil // if password is invalid this will return false otherwise it will return true
}