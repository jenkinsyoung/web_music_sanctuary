package hash

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func PasswordHash(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing occured %s", err)
	}
	return string(hashedPassword)
}

func CheckPassword(password, hashedPassword string) bool {
	//TODO: придумать как использовать эту функцию, с sql тут, или разбить на несколько функций?
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
