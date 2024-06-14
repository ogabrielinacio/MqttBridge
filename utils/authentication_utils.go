package utils

import (
	"crypto/hmac"
	"crypto/sha512"
	"math/rand"
	"regexp"
)

func CreatePasswordHash(password string) ([]byte, []byte, error) {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, nil, err
	}

	hasher := hmac.New(sha512.New, salt)
	hasher.Write([]byte(password))
	hash := hasher.Sum(nil)

	return hash, salt, nil
}

func IsSerialNumberValid(serialNumber string) bool {
	serialNumberRegex := "^[0-9]+-[0-9]+$"
	match, _ := regexp.MatchString(serialNumberRegex, serialNumber)
	return match
}
