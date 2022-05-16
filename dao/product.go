package dao

import (
	"echo-app/database"
	"echo-app/models"
	"echo-app/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListProduct(pipeline []bson.M) ([]models.ProductResponse, error) {
	var listProduct []models.ProductResponse
	cursor, err := database.ProductCol().Aggregate(utils.Ctx, pipeline)
	if err != nil {
		return listProduct, err
	}

	err = cursor.All(utils.Ctx, &listProduct)
	return listProduct, nil
}

func ProductDetail(id primitive.ObjectID) (*models.Product, error) {
	var results *models.Product
	err := database.ProductCol().FindOne(utils.Ctx, bson.M{"_id": id}).Decode(&results)

	return results, err
}
