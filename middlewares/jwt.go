package middlewares

import (
	"echo-app/config"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"strconv"
	"time"
)

type JwtCustomClaims struct {
	// ID String
	Data map[string]interface{}
	jwt.StandardClaims
}

var env = config.GetEnv()

func GenerateToken(data map[string]interface{}) (string, error) {
	// create expires time
	expTimeMs, err := strconv.Atoi(env.Jwt.TokenLife)
	if err != nil {
		fmt.Println(err)
	}
	exp := time.Now().Add(time.Millisecond * time.Duration(expTimeMs)).Unix()

	// init claims
	claims := JwtCustomClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	// generate token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err := token.SignedString([]byte(env.Jwt.SecretKey))

	return strToken, err
}

func GetJWTPayload(c echo.Context) (map[string]interface{}, error) {
	// get jwt object from context
	user := c.Get("user").(*jwt.Token)

	claims := &JwtCustomClaims{}

	// ParseWithClaims
	_, err := jwt.ParseWithClaims(user.Raw, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetEnv().Jwt.SecretKey), nil
	})

	// if err
	if err != nil {
		return nil, err
	}

	return claims.Data, nil
}
