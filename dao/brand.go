package dao

import (
	"context"
	"echo-app/database"
	"echo-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateBrand(brand models.BrandBSON) error {
	var (
		brandCol = database.BrandCol()
		ctx      = context.Background()
	)

	// InsertOne
	_, err := brandCol.InsertOne(ctx, brand)

	return err
}

func GetListBrand() ([]models.BrandBSON, error) {
	var (
		brandCol = database.BrandCol()
		ctx      = context.Background()
		brands   []models.BrandBSON
	)

	cursor, err := brandCol.Find(ctx, bson.D{})
	if err != nil {
		return brands, err
	}

	if err = cursor.All(ctx, &brands); err != nil {
		return brands, nil
	}

	return brands, nil
}

func GetBrandByID(ID primitive.ObjectID) (models.BrandBSON, error) {
	var (
		brandCol = database.BrandCol()
		ctx      = context.Background()
		brand    models.BrandBSON
	)

	// find brand
	filter := bson.M{"_id": ID}
	if err := brandCol.FindOne(ctx, filter).Decode(&brand); err != nil {
		return brand, err
	}

	return brand, nil
}

func UpdateBrandByID(ID primitive.ObjectID, body models.BrandUpdateBody) error {
	var (
		brandCol = database.BrandCol()
		ctx      = context.Background()
	)

	update := bson.M{"name": body.Name, "description": body.Description}

	// UpdateOne
	_, err := brandCol.UpdateOne(ctx, bson.M{"_id": ID}, bson.M{"$set": update})

	if err != nil {
		return err
	}

	return nil
}

//
func DeleteBrandByID(ID primitive.ObjectID) error {
	var (
		brandCol = database.BrandCol()
		ctx      = context.Background()
	)

	// filter
	filter := bson.M{"_id": ID}

	// DeleteOne
	if _, err := brandCol.DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}
