package services

import (
	"echo-app/dao"
	"echo-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"strings"
)

var limit int64 = 10

func ListProduct(query models.ProductQuery) ([]models.Product, error) {
	//init filter
	filter := bson.M{"page": query.Page}
	if query.CategoryId != "" {
		filter["CategoryId"], _ = primitive.ObjectIDFromHex(query.CategoryId)
	}
	if query.BrandId != "" {
		filter["BrandId"], _ = primitive.ObjectIDFromHex(query.BrandId)
	}
	if query.Name != "" {
		filter["Name"] = query.Name
	}
	if query.Name != "" {
		filter["Name"] = query.Name
	}

	// options query
	page, err := strconv.ParseInt(query.Page, 0, 64)
	if err != nil {
		return nil, err
	}
	optionsQuery := new(options.FindOptions)
	optionsQuery.SetSkip(page * limit)
	optionsQuery.SetLimit(limit)

	sort := strings.Split(query.Sort, ",")

	if query.Sort != "" {
		sortMap := map[string]interface{}{
			"price":     sort[0],
			"createdAt": sort[1],
		}
		optionsQuery.SetSort(sortMap)
	}

	results, err := dao.ListProduct(filter, optionsQuery)

	return results, err
}

func ProductDetail(id primitive.ObjectID) (model.Product, error) {
	return dao.ProductDetail(id), nil
}
