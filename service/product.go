package service

import (
	"echo-app/dao"
	"echo-app/model"
	"fmt"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var limit int64 = 10

type Product struct{}

func (p Product) ListProduct(query model.ProductQuery) ([]model.ProductResponse, error) {
	var (
		d      = dao.Product{}
		sort   = bson.M{}
		filter = bson.M{"status": "enable"}
	)

	//init filter
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

	//init sort
	if query.Sort != "" {
		sortArr := strings.Split(query.Sort, ",")
		for _, v := range sortArr {
			var value int
			if string([]rune(v)[1]) != "-" {
				fmt.Println(string([]rune(v)[0]))
				value = -1
			} else {
				value = 1
			}
			sort[v] = value
		}
	}

	results, err := d.ListProduct(filter, query, sort)

	return results, err
}

func (p Product) ProductDetail(id string) (*model.ProductResponse, error) {
	ojbId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	results, err := productDAO.ProductDetail(ojbId)
	return results, err
}
