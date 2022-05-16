package services

import (
	"echo-app/dao"
	"echo-app/models"
	"echo-app/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ListReview(productId primitive.ObjectID, query models.ReviewQuery) ([]models.Review, error) {
	filter := bson.M{"productId": productId, "rating": query.Rating}

	optionsQuery := new(options.FindOptions)
	optionsQuery.SetSkip(query.Page * limit)
	optionsQuery.SetLimit(limit)
	if query.Sort != "" {
		sortMap := map[string]interface{}{
			"price": query.Sort,
		}
		optionsQuery.SetSort(sortMap)
	}

	results, err := dao.ListReview(filter, optionsQuery)
	return results, err
}

func CreateReview(userId primitive.ObjectID, productId primitive.ObjectID, body models.CreateReview) error {
	//init insert data
	insertData := models.Review{
		ID:        primitive.NewObjectID(),
		UserId:    userId,
		ProductId: productId,
		Rating:    body.Rating,
		Content:   body.Content,
		CreatedAt: utils.CurrentDateTime(),
		UpdatedAt: utils.CurrentDateTime(),
	}
	err := dao.CreateReview(insertData)
	return err
}
