package dao

import (
	"echo-app/database"
	"echo-app/models"
	"echo-app/utils"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ListProduct(filter interface{}, options *options.FindOptions) ([]models.Product, error) {
	var listProduct []models.Product
	cursor, err := database.ProductCol().Find(utils.Ctx, filter, options)
	if err != nil {
		fmt.Println(err.Error())
		return listProduct, err
	}

	if err := cursor.All(utils.Ctx, &listProduct); err != nil {
		return listProduct, err
	}
	return listProduct, nil
}

func ProductDetail(id primitive.ObjectID) (models.Product, error) {
	var results models.Product
	if err := database.ProductCol().FindOne(utils.Ctx, bson.M{"id": id}).Decode(&results); err != nil {
		return results, err
	}

	return results, nil
}
