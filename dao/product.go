package dao

import (
	"echo-app/database"
	"echo-app/models"
	"echo-app/utils"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListProduct(pipeline []bson.M) ([]models.ProductResponse, error) {
	var listProduct []models.ProductResponse
	//cursor, err := database.ProductCol().Find(utils.Ctx, filter, options)
	cursor, err := database.ProductCol().Aggregate(utils.Ctx, pipeline)
	if err != nil {
		return listProduct, err
	}

	var results []bson.M
	err = cursor.All(utils.Ctx, &listProduct)
	fmt.Println(results)
	fmt.Println(listProduct)
	//bsonBytes, err := json.Marshal(results)
	//if err != nil {
	//	return listProduct, nil
	//}
	//bson.Unmarshal(bsonBytes, &listProduct)
	return listProduct, nil
}

func ProductDetail(id primitive.ObjectID) (*models.Product, error) {
	var results *models.Product
	err := database.ProductCol().FindOne(utils.Ctx, bson.M{"_id": id}).Decode(&results)

	return results, err
}
