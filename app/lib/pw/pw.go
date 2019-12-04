package pw

import (
	"golang.org/x/crypto/bcrypt"
)

// GenerateHash builds a bcrypted storable password
func GenerateHash(pwd string) string {

	password := []byte(pwd)

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

// IsPasswordValid compares the hash with the test pwd
func IsPasswordValid(hashedPassword string, pwd string) bool {
	hashBytes := []byte(hashedPassword)
	pwdBytes := []byte(pwd)
	err := bcrypt.CompareHashAndPassword(hashBytes, pwdBytes)
	if err != nil {
		panic(err)
	}
	return true
}
