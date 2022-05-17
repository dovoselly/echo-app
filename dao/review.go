package dao

import (
	"context"
	"echo-app/database"
	"echo-app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ListReview(filter bson.M, options *options.FindOptions) ([]models.Review, error) {
	var (
		ctx = context.Background()
	)

	var listReview []models.Review
	cursor, err := database.ReviewCol().Find(ctx, filter, options)
	if err != nil {
		return listReview, err
	}

	err = cursor.All(ctx, &listReview)
	return listReview, err
}

func CreateReview(insertData models.Review) error {
	var (
		ctx = context.Background()
	)

	_, err := database.ReviewCol().InsertOne(ctx, insertData)
	if err != nil {
		return err
	}

	return nil
}
