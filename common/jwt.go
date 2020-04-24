package common

import (
	"WebFull/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT中间件认证
var jwtKey = []byte("a_secret_cred")

type Claim struct {
	jwt.StandardClaims
	UserId uint
}

func ReleaseToken(u model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claim{
		UserId: u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "yang_zhang_gin_supermarket",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenStr string) (*jwt.Token, *Claim, error) {
	claim := &Claim{}
	token, err := jwt.ParseWithClaims(tokenStr, claim, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claim, err
}
