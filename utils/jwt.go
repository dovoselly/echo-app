package utils

import (
	"echo-app/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// JwtCustomClaims ...
type JwtCustomClaims struct {
	//ID string
	Data map[string]interface{}
	jwt.StandardClaims
}

// GenerateUserToken ...
func GenerateUserToken(data map[string]interface{}) (string, error) {

	// claims ...
	claims := &JwtCustomClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// generate token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token
	str, err := token.SignedString([]byte(config.GetEnv().Jwt.SecretKey))

	// if err
	if err != nil {
		return "", err
	}

	return str, nil
}

// GetJWTPayload ...
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

func GetIdInToken(c echo.Context) (primitive.ObjectID, error) {

	// GetJWTPaylaod
	jwtPayload, err := GetJWTPayload(c)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	idString := jwtPayload["id"].(string)

	ID, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return ID, nil
}
