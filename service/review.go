package service

import (
	"echo-app/model"
	"go.mongodb.org/mongo-driver/mongo"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct{}

func (Review) GetListReview(ID string, query model.ReviewQuery) ([]model.ReviewResponse, error) {
	ojbID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return []model.ReviewResponse{}, err
	}

	results, err := reviewDAO.GetListReview(ojbID, query)
	return results, err
}

func (Review) CreateReview(userId primitive.ObjectID, productId primitive.ObjectID, body model.CreateReview) (*mongo.InsertOneResult, error) {
	//init insert data
	insertData := model.ReviewBSON{
		ID:        primitive.NewObjectID(),
		UserId:    userId,
		ProductId: productId,
		Rating:    body.Rating,
		Content:   body.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result, err := reviewDAO.CreateReview(insertData)
	return result, err
}
