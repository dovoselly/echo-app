package dao

import (
	"echo-app/database"
	"echo-app/model"
	"echo-app/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Brand struct{}

func (Brand) CreateBrand(brand model.BrandBSON) (string, error) {
	// InsertOne
	result, err := database.BrandCol().InsertOne(util.Ctx, brand)

	return result.InsertedID.(string), err
}

func (Brand) GetListBrand() ([]model.BrandBSON, error) {
	var brands []model.BrandBSON

	cursor, err := database.BrandCol().Find(util.Ctx, bson.D{})
	if err != nil {
		return brands, err
	}

	if err = cursor.All(util.Ctx, &brands); err != nil {
		return brands, nil
	}

	return brands, nil
}

func (Brand) GetBrandByID(ID primitive.ObjectID) (model.BrandBSON, error) {
	var (
		brand model.BrandBSON
	)

	// find brand
	filter := bson.M{"_id": ID}
	if err := database.BrandCol().FindOne(util.Ctx, filter).Decode(&brand); err != nil {
		return brand, err
	}

	return brand, nil
}

func (Brand) UpdateBrandByID(ID primitive.ObjectID, body model.BrandUpdateBody) error {
	update := bson.M{"name": body.Name, "description": body.Description}

	// UpdateOne
	_, err := database.BrandCol().UpdateOne(util.Ctx, bson.M{"_id": ID}, bson.M{"$set": update})

	if err != nil {
		return err
	}

	return nil
}

func (Brand) DeleteBrandByID(ID primitive.ObjectID) error {
	// filter
	filter := bson.M{"_id": ID}

	// DeleteOne
	if _, err := database.BrandCol().DeleteOne(util.Ctx, filter); err != nil {
		return err
	}

	return nil
}
