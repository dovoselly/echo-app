package dao

import (
	"echo-app/database"
	"echo-app/model"
	"echo-app/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct{}

func (p Product) ListProduct(pipeline []bson.M) ([]model.ProductResponse, error) {
	var listProduct []model.ProductResponse
	cursor, err := database.ProductCol().Aggregate(utils.Ctx, pipeline)
	if err != nil {
		return listProduct, err
	}

	err = cursor.All(utils.Ctx, &listProduct)
	return listProduct, nil
}

func (p Product) ProductDetail(id primitive.ObjectID) (*model.ProductResponse, error) {

	var results *model.ProductResponse
	err := database.ProductCol().FindOne(utils.Ctx, bson.M{"_id": id}).Decode(&results)

	return results, err
}
