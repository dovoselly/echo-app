package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	UserBSON struct {
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
		CreatedAt      time.Time          `bson:"createdAt"`
		UssspdatedAt   time.Time          `bson:"updatedAt"`
	}

	UserResponse struct {
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
		CreatedAt      time.Time          `json:"createdAt"`
		UssspdatedAt   time.Time          `json:"updatedAt"`
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

	UserInfo struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)
