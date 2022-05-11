package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Admin struct {
		ID             primitive.ObjectID `bson:"_id"`
		Email          string             `bson:"email"`
		Username       string             `bson:"username"`
		HashedPassword string             `bson:"hashedPassword"`
		FullName       string             `bson:"fullName"`
		DateOfBirth    string             `bson:"dateOfBirth"`
		Avatar         string             `bson:"avatar"`
		Gender         string             `bson:"gender"`
		Phone          string             `bson:"phone"`
		CreatedAt      string             `bson:"createdAt"`
		UpdatedAt      string             `bson:"updatedAt"`
	}

	// AdminLogin
	AdminLoginBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

func (a AdminLoginBody) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Username, validation.Required),
		validation.Field(&a.Password, validation.Required),
	)
}
