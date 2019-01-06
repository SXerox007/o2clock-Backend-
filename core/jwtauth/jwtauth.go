package jwtauth

import (
	"crypto/rsa"
	"time"

	"o2clock/constants/collections"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	EXP = "exp"
	SUB = "sub"
)

//Generate the Access token
func GenerateToken(key *rsa.PrivateKey, username string, uuid string) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodRS512)
	claims := make(jwt.MapClaims)
	claims[collections.PARAM_USER_NAME] = username
	claims[EXP] = time.Now().Add(time.Minute * 1).Unix()
	claims[SUB] = uuid
	token.Claims = claims
	return token.SignedString(key)
}

//Logout the user
func Logout(tokenString string, token *jwt.Token) error {
	return nil
}

//Blacklist the token can't access
func BlackListToken(tokenString string) error {
	return nil
}

//parse the token
func ParseToken(myToken string, key []byte) bool {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err == nil && token.Valid {
		return true
	} else {
		return false
	}
}
