package dao

import (
	"echo-app/database"
	"echo-app/model"
	"echo-app/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ListReview(filter bson.M, options *options.FindOptions) ([]model.Review, error) {
	var listReview []model.Review
	cursor, err := database.ReviewCol().Find(utils.Ctx, filter, options)
	if err != nil {
		return listReview, err
	}

	err = cursor.All(utils.Ctx, &listReview)
	return listReview, err
}

func CreateReview(insertData model.Review) error {
	_, err := database.ReviewCol().InsertOne(utils.Ctx, insertData)
	if err != nil {
		return err
	}
	return nil
}
