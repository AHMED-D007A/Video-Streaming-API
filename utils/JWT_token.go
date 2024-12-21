package utils

import (
	"Video-Streaming-API/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const EmailKey contextKey = "email"

type CustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateToken(Name string, Email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaims{
		Email: Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	})

	tokenStr, err := token.SignedString([]byte(config.Envs.JWT_SECRET))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

/*
	 A function that takes the token string and parse it to get the claims
		and return the email field of the custom claims and err.
*/
func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		alg := t.Method.Alg()
		if alg != jwt.SigningMethodHS256.Name {
			return nil, errors.New("invalid Signing Method")
		}

		return config.Envs.JWT_SECRET, nil
	})
	if err != nil {
		return "", err
	}

	claims, _ := token.Claims.(*CustomClaims)

	return claims.Email, nil
}
