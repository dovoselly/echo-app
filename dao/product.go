package dao

import (
	"echo-app/database"
	"echo-app/models"
	"echo-app/utils"
	"fmt"
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

func ProductDetails(id primitive.ObjectID) (models.Product, error) {
	results := database.ProductCol().FindOne()
}
