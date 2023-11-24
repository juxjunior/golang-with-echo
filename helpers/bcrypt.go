package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPass(p string) string {
	salt := 0
	pass := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(pass, salt)
	if err != nil {
		panic(err)
	}

	return string(hash)
}

func ComparePass(h, p []byte) bool {
	err := bcrypt.CompareHashAndPassword(h, p)
	if err != nil {
		return false
	}
	return true
}
