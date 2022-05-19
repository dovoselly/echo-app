package service

import (
	"echo-app/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Brand struct{}

func (b Brand) CreateBrand(body model.BrandCreateBody) (string, error) {
	// brand BSON

	brand := model.BrandBSON{
		ID:          primitive.NewObjectID(),
		Name:        body.Name,
		Description: body.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// create brand
	insertedID, err := brandDAO.CreateBrand(brand)
	if err != nil {
		return "", err
	}

	return insertedID, nil
}

func (b Brand) GetList() ([]model.BrandResponse, error) {

	listBrand := make([]model.BrandResponse, 0)

	// get list brand bson
	brandsBSON, err := brandDAO.GetListBrand()
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
	brandBSON, err := brandDAO.GetBrandByID(objID)

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

	err := brandDAO.UpdateBrandByID(objID, body)
	if err != nil {
		return err
	}
	return nil

}

func (b Brand) DeleteByID(ID string) error {
	// convert id string to objectID
	objID, _ := primitive.ObjectIDFromHex(ID)

	err := brandDAO.DeleteBrandByID(objID)
	if err != nil {
		return err
	}

	// success
	return nil
}
