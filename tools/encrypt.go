package tools

import (
	"golang.org/x/crypto/bcrypt"
)

func PasswordEncrypt(password string) string {
	encrypted, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(encrypted)
}

func PasswordCompare(password, encrypted string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(password))
	return err == nil
}
