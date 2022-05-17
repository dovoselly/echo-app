package dao

import (
	"echo-app/database"
	"echo-app/model"
	"echo-app/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListProduct(pipeline []bson.M) ([]model.ProductResponse, error) {
	var listProduct []model.ProductResponse
	cursor, err := database.ProductCol().Aggregate(util.CONTEXT, pipeline)
	if err != nil {
		return listProduct, err
	}

	err = cursor.All(util.CONTEXT, &listProduct)
	return listProduct, nil
}

func ProductDetail(id primitive.ObjectID) (*model.ProductBSON, error) {
	var results *model.ProductBSON
	err := database.ProductCol().FindOne(util.CONTEXT, bson.M{"_id": id}).Decode(&results)

	return results, err
}
