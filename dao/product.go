package dao

import (
	"context"
	"echo-app/database"
	"echo-app/models"
	"echo-app/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListProduct(pipeline []bson.M) ([]models.ProductResponse, error) {
	var ()

	var listProduct []models.ProductResponse
	cursor, err := database.ProductCol().Aggregate(utils.CONTEXT, pipeline)
	if err != nil {
		return listProduct, err
	}

	err = cursor.All(utils.CONTEXT, &listProduct)
	return listProduct, nil
}

func ProductDetail(id primitive.ObjectID) (*models.Product, error) {
	var (
		ctx = context.Background()
	)

	var results *models.Product
	err := database.ProductCol().FindOne(ctx, bson.M{"_id": id}).Decode(&results)

	return results, err
}
