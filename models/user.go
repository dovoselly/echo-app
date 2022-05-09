package models

import "go.mongodb.org/mongo-driver/bson/primitive"

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
