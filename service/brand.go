package service

import (
	"echo-app/dao"
	"echo-app/model"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateBrand(brandBody model.BrandCreateBody) error {
	// category BSON

	brand := model.BrandBSON{
		ID:          primitive.NewObjectID(),
		Name:        brandBody.Name,
		Description: brandBody.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// create brand
	if err := dao.CreateBrand(brand); err != nil {
		return errors.New("can not create new brand")
	}

	return nil
}

func GetListBrand() ([]model.BrandResponse, error) {

	listBrand := make([]model.BrandResponse, 0)

	// get list brand bson
	brandsBSON, err := dao.GetListBrand()
	if err != nil {
		return listBrand, err
	}

	for _, brandBSON := range brandsBSON {
		brandJSON := model.BrandResponse{
			ID:          brandBSON.ID,
			Name:        brandBSON.Name,
			Description: brandBSON.Description,
			Status:      brandBSON.Status,
		}
		listBrand = append(listBrand, brandJSON)
	}

	return listBrand, nil

}

func GetBrandByID(ID string) (model.BrandResponse, error) {
	var (
		brand model.BrandResponse
	)

	// to objectID
	objID, _ := primitive.ObjectIDFromHex(ID)

	// get brand by id
	brandBSON, err := dao.GetBrandByID(objID)

	brand = model.BrandResponse{
		ID:          brandBSON.ID,
		Name:        brandBSON.Name,
		Description: brandBSON.Description,
		Status:      brandBSON.Status,
		CreatedAt:   brandBSON.CreatedAt,
		UpdatedAt:   brandBSON.UpdatedAt,
	}

	if err != nil {
		return brand, err
	}

	return brand, nil
}

func UpdateBrandByID(ID string, body model.BrandUpdateBody) error {
	objID, _ := primitive.ObjectIDFromHex(ID)

	err := dao.UpdateBrandByID(objID, body)
	if err != nil {
		return err
	}
	return nil

}

func DeleteBrandByID(ID string) error {
	// convert id string to objectID
	objID, _ := primitive.ObjectIDFromHex(ID)

	err := dao.DeleteBrandByID(objID)
	if err != nil {
		return err
	}

	// success
	return nil
}
