package controllers

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)


// Create a new token object, specifying signing method and the claims
// you would like it to contain.


func GenerateAccessToken(userId int) string {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(30 * time.Minute).Unix(),
	})

	// Convert the signing key to []byte
	key := []byte("secret")

	// Sign and get the complete encoded token as a string using the key
	tokenString, err := accessToken.SignedString(key)

	if err != nil {
		return err.Error()
	}

	return tokenString
	
}

func GenerateRefreshToken(userId int) string {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	key := []byte("secret")

	tokenString, err := refreshToken.SignedString(key)

	if err != nil {
		return err.Error()
	}

	return tokenString
}