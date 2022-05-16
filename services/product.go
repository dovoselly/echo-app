package services

import (
	"echo-app/dao"
	"echo-app/models"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"strings"
)

var limit int64 = 10

func ListProduct(query models.ProductQuery) ([]models.ProductResponse, error) {
	//init filter
	filter := bson.M{"status": "enable"}

	if query.CategoryId != "" {
		filter["categoryId"], _ = primitive.ObjectIDFromHex(query.CategoryId)
	}
	if query.BrandId != "" {
		filter["brandId"], _ = primitive.ObjectIDFromHex(query.BrandId)
	}
	if query.Name != "" {
		filter["name"] = query.Name
	}
	if query.PriceFrom != "" {
		priceFrom, _ := strconv.ParseInt(query.PriceFrom, 10, 64)
		filter["price"] = bson.M{
			"$lte": priceFrom,
		}
	}

	// options query
	//optionsQuery := new(options.FindOptions)
	//optionsQuery.SetSkip(query.Page * limit)
	//optionsQuery.SetLimit(limit)

	pipeline := []bson.M{
		{"$match": filter},
		{"$lookup": bson.M{
			"from":         "brands",
			"localField":   "brandId",
			"foreignField": "_id",
			"as":           "brand",
		}},
		{"$lookup": bson.M{
			"from":         "categories",
			"localField":   "categoryId",
			"foreignField": "_id",
			"as":           "category",
		}},
		{"$skip": query.Page * limit},
		{"$limit": limit},
		{"$unwind": "$category"},
		{"$unwind": "$brand"},
	}

	if query.Sort != "" {
		sortArr := strings.Split(query.Sort, ",")
		sortMap := bson.M{}
		for _, v := range sortArr {
			var value int
			if string([]rune(v)[1]) != "-" {
				fmt.Println(string([]rune(v)[0]))
				value = -1
			} else {
				value = 1
			}
			sortMap[v] = value
		}
		pipeline = append(pipeline, bson.M{"$sort": sortMap})
	}

	results, err := dao.ListProduct(pipeline)

	return results, err
}

func ProductDetail(id primitive.ObjectID) (*models.Product, error) {
	results, err := dao.ProductDetail(id)
	return results, err
}
