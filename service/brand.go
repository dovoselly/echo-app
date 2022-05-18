package service

import (
	"echo-app/model"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Brand struct{}

func (b Brand) Create(brandBody model.BrandCreateBody) error {
	// brand BSON

	brand := model.BrandBSON{
		ID:          primitive.NewObjectID(),
		Name:        brandBody.Name,
		Description: brandBody.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// create brand
	if err := brandDao.Create(brand); err != nil {
		return errors.New("can not create new brand")
	}

	return nil
}

func (b Brand) GetList() ([]model.BrandResponse, error) {

	listBrand := make([]model.BrandResponse, 0)

	// get list brand bson
	brandsBSON, err := brandDao.GetList()
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

func (b Brand) GetByID(ID string) (model.BrandResponse, error) {
	var (
		brand model.BrandResponse
	)

	// to objectID
	objID, _ := primitive.ObjectIDFromHex(ID)

	// get brand by id
	brandBSON, err := brandDao.GetByID(objID)

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

func (b Brand) UpdateByID(ID string, body model.BrandUpdateBody) error {
	objID, _ := primitive.ObjectIDFromHex(ID)

	err := brandDao.UpdateByID(objID, body)
	if err != nil {
		return err
	}
	return nil

}

func (b Brand) DeleteByID(ID string) error {
	// convert id string to objectID
	objID, _ := primitive.ObjectIDFromHex(ID)

	err := brandDao.DeleteByID(objID)
	if err != nil {
		return err
	}

	// success
	return nil
}
