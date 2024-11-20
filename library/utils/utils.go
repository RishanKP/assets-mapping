package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func IsEmpty(str string) bool {
	if len(str) == 0 {
		return true
	}

	return false
}

func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPassword), err
}
