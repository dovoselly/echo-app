package utils

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserId(c echo.Context) (primitive.ObjectID, error) {
	// GetJWTPayload
	jwtPayload, err := GetJWTPayload(c)
	if err != nil {
		return primitive.NilObjectID, err
	}

	idString, ok := jwtPayload["id"].(string)
	if !ok {
		return primitive.NilObjectID, err
	}

	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return id, nil
}
