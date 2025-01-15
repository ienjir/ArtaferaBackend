package auth

import (
	_ "github.com/golang-jwt/jwt/v5"
)

var Argon2IDHash *Argon2idHash

func HashPassword(password string) (*HashSalt, error) {
	bytePassword := []byte(password)

	hashSalt, err := Argon2IDHash.GenerateHash(bytePassword, nil)
	if err != nil {
		return nil, err
	}

	return hashSalt, nil
}
