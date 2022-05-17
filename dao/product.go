package dao

import (
	"echo-app/database"
	"echo-app/model"
	"echo-app/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var limit int64 = 10

type Product struct{}

func (p Product) GetListProduct(filter bson.M, query model.ProductQuery, sort bson.M) ([]model.ProductResponse, error) {
	var (
		listProduct []model.ProductResponse
		pipeline    = []bson.M{
			{"$match": filter},
			{"$lookup": bson.M{
				"from":         "brands",
				"localField":   "brandId",
				"foreignField": "_id",
				"as":           "brand",
			}},
			{"$unwind": "$brand"},
			{"$lookup": bson.M{
				"from":         "categories",
				"localField":   "categoryId",
				"foreignField": "_id",
				"as":           "category",
			}},
			{"$unwind": "$category"},
			{"$skip": query.Page * limit},
			{"$limit": limit},
			{"sort": sort},
		}
	)
	cursor, err := database.ProductCol().Aggregate(utils.Ctx, pipeline)
	if err != nil {
		return listProduct, err
	}

	err = cursor.All(utils.Ctx, &listProduct)
	return listProduct, nil
}

func (p Product) GetProductDetail(id primitive.ObjectID) (*model.ProductResponse, error) {
	var (
		results *model.ProductResponse
		filter  = bson.M{"_id": id}
	)

	err := database.ProductCol().FindOne(utils.Ctx, filter).Decode(&results)

	return results, err
}
