package services

import (
	"echo-app/dao"
	"echo-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

var limit int64 = 10

func ListProduct(query models.ProductQuery) ([]models.Product, error) {
	//init filter
	filter := bson.M{}
	if query.CategoryId != "" {
		filter["CategoryId"], _ = primitive.ObjectIDFromHex(query.CategoryId)
	}
	if query.BrandId != "" {
		filter["BrandId"], _ = primitive.ObjectIDFromHex(query.BrandId)
	}
	if query.Name != "" {
		filter["Name"] = query.Name
	}
	if query.PriceFromTo != "" {
		priceFromTo := strings.Split(query.PriceFromTo, ",")
		filter["PriceFromTo"] = bson.M{"$and": []bson.M{
			{"price": bson.M{"$gte": priceFromTo[0]}},
			{"price": bson.M{"$lte": priceFromTo[1]}},
		},
		}
	}

	// options query
	optionsQuery := new(options.FindOptions)
	optionsQuery.SetSkip(query.Page * limit)
	optionsQuery.SetLimit(limit)

	if query.Sort != "" {
		sort := strings.Split(query.Sort, ",")
		sortMap := map[string]interface{}{
			"price":     sort[0],
			"createdAt": sort[1],
		}
		optionsQuery.SetSort(sortMap)
	}

	results, err := dao.ListProduct(filter, optionsQuery)

	return results, err
}

func ProductDetail(id primitive.ObjectID) (*models.Product, error) {
	results, err := dao.ProductDetail(id)
	return results, err
}
