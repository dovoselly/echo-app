package utils

import (
	"echo-app/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JwtCustomClaims ...
type JwtCustomClaims struct {
	//ID string
	Data map[string]interface{}
	jwt.StandardClaims
}

var envVars = config.GetEnv()

// GenerateUserToken ...
func GenerateUserToken(data map[string]interface{}) (string, error) {

	// claims ...
	claims := &JwtCustomClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 120).Unix(),
		},
	}

	// generate token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token
	str, err := token.SignedString([]byte(envVars.Jwt.SecretKey))

	// if err
	if err != nil {
		return "", err
	}

	return str, nil
}
