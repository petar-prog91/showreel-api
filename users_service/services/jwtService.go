package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

var tokenSigningKey = []byte("y42jh9824j9h82j49h82j40g9im240h9240h94p2hjk0249h0924")

type MyJWTClaims struct {
	UserName string `json:"userName"`
	UserRole int `json:"userRole"`
	jwt.StandardClaims
}

func GenerateNewToken(userName string, userRole int) (string, error) {
	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
		"userName": userName,
		"userRole": userRole,
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(tokenSigningKey)

	return tokenString, err
}

func ParseToken(myToken string) (bool, *MyJWTClaims, error) {
	token, err := jwt.ParseWithClaims(myToken, &MyJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return tokenSigningKey, nil
	})

	if claims, ok := token.Claims.(*MyJWTClaims); ok && token.Valid {
		return true, claims, nil
	} else {
		fmt.Println(err)
		return false, nil, err
	}
}
