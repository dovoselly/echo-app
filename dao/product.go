package dao

import (
	"echo-app/database"
	"echo-app/model"
	"echo-app/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListProduct(pipeline []bson.M) ([]model.ProductResponse, error) {
	var listProduct []model.ProductResponse
	cursor, err := database.ProductCol().Aggregate(utils.Ctx, pipeline)
	if err != nil {
		return listProduct, err
	}

	err = cursor.All(utils.CONTEXT, &listProduct)
	return listProduct, nil
}

func ProductDetail(id primitive.ObjectID) (*model.Product, error) {
	var results *model.Product
	err := database.ProductCol().FindOne(utils.Ctx, bson.M{"_id": id}).Decode(&results)

	return results, err
}
