package dao

import (
	"echo-app/database"
	"echo-app/model"
	"echo-app/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ListReview(filter bson.M, options *options.FindOptions) ([]model.Review, error) {
	var listReview []model.Review
	cursor, err := database.ReviewCol().Find(util.CONTEXT, filter, options)
	if err != nil {
		return listReview, err
	}

	err = cursor.All(util.CONTEXT, &listReview)
	return listReview, err
}

func CreateReview(insertData model.Review) error {
	_, err := database.ReviewCol().InsertOne(util.CONTEXT, insertData)
	if err != nil {
		return err
	}
	return nil
}
