package dao

import (
	"echo-app/database"
	"echo-app/model"
	"echo-app/util"
)

type Cart struct{}

func (c Cart) Create(body model.CartCreateBSON) (string, error) {
	result, err := database.CartCol().InsertOne(util.Ctx, body)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(string), err
}
