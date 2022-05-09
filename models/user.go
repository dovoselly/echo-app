package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
)

type (
	User struct {
		ID             primitive.ObjectID `bson:"_id"`
		Email          string             `bson:"email" `
		Username       string             `bson:"username"`
		HashedPassword string             `bson:"hashedPassword"`
		Name           string             `bson:"name"`
		DateOfBirth    string             `bson:"dateOfBirth"`
		Avatar         string             `bson:"avatar"`
		Gender         string             `bson:"gender"`
		Phone          string             `bson:"phone"`
		Status         string             `bson:"status"`
		CreatedAt      string             `bson:"createdAt"`
		UpdatedAt      string             `bson:"updatedAt"`
	}

	UserDetail struct {
		ID             primitive.ObjectID `json:"_id"`
		Email          string             `json:"email"`
		Username       string             `json:"username"`
		HashedPassword string             `json:"hashedPassword"`
		Name           string             `json:"name"`
		DateOfBirth    string             `json:"dateOfBirth"`
		Avatar         string             `json:"avatar"`
		Gender         string             `json:"gender"`
		Phone          string             `json:"phone"`
		Status         string             `json:"status"`
		CreatedAt      string             `json:"createdAt"`
		UpdatedAt      string             `json:"updatedAt"`
	}

	UserRegister struct {
		Email       string `json:"email"`
		Username    string `json:"userName"`
		Password    string `json:"password"`
		Name        string `json:"name"`
		DateOfBirth string `json:"dateOfBirth"`
		Gender      string `json:"gender"`
	}

	UserLogin struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
	}

	UserChangePassword struct {
		CurrentPassword string `json:"currentPassword"`
		NewPassword     string `json:"newPassword"`
	}

	UserUpdate struct {
		FullName    string `json:"fullName"`
		Email       string `json:"email"`
		Phone       string `json:"phone"`
		DateOfBirth string `json:"dateOfBirth"`
		Gender      string `json:"gender"`
		Address     string `json:"address"`
	}
)

func (u UserRegister) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Match(regexp.MustCompile("^[a-zA-Z0-9_]{8,64}$"))),
		validation.Field(&u.Username, validation.Match(regexp.MustCompile("^[a-zA-Z0-9_]{8,64}$"))),
		validation.Field(&u.Password, validation.Match(regexp.MustCompile("^[a-zA-Z0-9_]{8,64}$"))),
	)
}
