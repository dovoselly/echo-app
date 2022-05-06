package middleware

import (
	"echo-app/config"
	"echo-app/util"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Claims struct {
	ID string `json:"_id" bson:"_id"`
	jwt.StandardClaims
}

func GenerateToken(id primitive.ObjectID) (string, error) {
	env := config.GetEnv()

	// create expires time
	expTimeMs, err := strconv.Atoi(env.Jwt.TokenLife)
	if err != nil {
		fmt.Println(err)
	}
	exp := time.Now().Add(time.Millisecond * time.Duration(expTimeMs)).Unix()

	// init claims
	claims := Claims{
		id.Hex(),
		jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	// generate token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err := token.SignedString([]byte(env.Jwt.SecretKey))

	return strToken, err
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		env := config.GetEnv()
		// replace Bearer Token
		token := strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", 1)

		//parse Token
		parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(env.Jwt.SecretKey), nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, util.Response{
				Message: "Invalid Token",
			})
		}

		// validate token and set id token
		if claims, ok := parsedToken.Claims.(*Claims); ok && parsedToken.Valid {
			c.Set("id", claims.ID)
			return next(c)
		} else {
			return c.JSON(http.StatusUnauthorized, util.Response{
				Message: util.InvalidToken,
			})
		}
	}
}
