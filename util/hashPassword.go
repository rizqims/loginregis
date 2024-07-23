package util

import (
	"crypto/sha256"
	"fmt"
)

func HashPassword(payloadPass string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(payloadPass))
	if err != nil {
		return "", fmt.Errorf("hashPassword error: %v", err)
	}
	hashBytes := hash.Sum(nil)
	hashString := fmt.Sprintf("%x", hashBytes)
	fmt.Println(hashString)

	return hashString, nil
}

func CheckHash(payloadPass string, userPass string) error {
	hashedPass, err := HashPassword(payloadPass)
	if err != nil {
		return fmt.Errorf("checkHash error: hashPassword error: %v", err)
	}
	if hashedPass != userPass {
		return fmt.Errorf("CheckHash error: password is invalid!")
	}
	return nil
}
