package service

import (
	"echo-app/dao"
	"echo-app/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var limit int64 = 10

type Product struct{}

func (p Product) GetListProduct(query model.ProductQuery) ([]model.ProductResponse, error) {
	var d = dao.Product{}
	results, err := d.GetListProduct(query)
	return results, err
}

func (p Product) GetProductDetail(id string) (*model.ProductResponse, error) {
	ojbId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	results, err := productDAO.GetProductDetail(ojbId)
	return results, err
}
