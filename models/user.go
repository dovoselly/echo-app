package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	UserBSON struct {
		ID          primitive.ObjectID `bson:"_id"`
		Email       string             `bson:"email" `
		Username    string             `bson:"username"`
		Password    string             `bson:"password"`
		FullName    string             `bson:"fullName"`
		DateOfBirth string             `bson:"dateOfBirth"`
		Avatar      string             `bson:"avatar"`
		Gender      string             `bson:"gender"`
		Phone       string             `bson:"phone"`
		Address     string             `bson:"address"`
		Status      string             `bson:"status"`
		CreatedAt   time.Time          `bson:"createdAt"`
		UpdatedAt   time.Time          `bson:"updatedAt"`
	}

	UserInfoBSON struct {
		FullName    string `bson:"fullName"`
		Email       string `bson:"email" `
		DateOfBirth string `bson:"dateOfBirth"`
		Gender      string `bson:"gender"`
		Phone       string `bson:"phone"`
		Address     string `bson:"address"`
		Status      string `bson:"status"`
	}

	UserResponse struct {
		ID          primitive.ObjectID `json:"_id"`
		Email       string             `json:"email"`
		Username    string             `json:"username"`
		Password    string             `json:"password"`
		FullName    string             `json:"fullName"`
		DateOfBirth string             `json:"dateOfBirth"`
		Avatar      string             `json:"avatar"`
		Gender      string             `json:"gender"`
		Phone       string             `json:"phone"`
		Address     string             `bson:"address"`
		Status      string             `json:"status"`
		CreatedAt   time.Time          `json:"createdAt"`
		UpdatedAt   time.Time          `json:"updatedAt"`
	}

	UserLogin struct {
		Username string `json:"username"`
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

	UserRegister struct {
		Username    string `json:"username"`
		Password    string `json:"password"`
		Email       string `json:"email"`
		FullName    string `json:"fullName"`
		Gender      string `json:"gender"`
		DateOfBirth string `json:"dateOfBirth"`
		Phone       string `json:"phone"`
		Address     string `json:"address"`
		Status      string `json:"status"`
	}

	UserInfo struct {
		ID          primitive.ObjectID `json:"_id"`
		FullName    string             `json:"fullName"`
		Email       string             `json:"email"`
		Username    string             `json:"username"`
		DateOfBirth string             `json:"dateOfBirth"`
		Avatar      string             `json:"avatar"`
		Gender      string             `json:"gender"`
		Phone       string             `json:"phone"`
		Address     string             `json:"address"`
	}
)

// ConvertToBSON
func (b UserRegister) ConvertToBSON() UserBSON {
	result := UserBSON{
		ID:          primitive.NewObjectID(),
		Username:    b.Username,
		Password:    b.Password,
		Email:       b.Email,
		FullName:    b.FullName,
		Gender:      b.Gender,
		DateOfBirth: b.DateOfBirth,
		Phone:       b.Phone,
		Address:     b.Address,
		Status:      b.Status,
		CreatedAt:   time.Now(),
	}
	return result
}

// Validate form body
func (payload UserRegister) Validate() error {
	return validation.ValidateStruct(&payload,

		validation.Field(
			&payload.Username,
			validation.Required.Error("Username is required"),
			validation.Length(5, 30).Error("UserName is length: 5 -> 30"),
		),

		validation.Field(
			&payload.Password,
			validation.Required.Error("Password is required"),
			validation.Length(5, 30).Error("Password is length: 5 -> 30"),
		),

		validation.Field(
			&payload.Email,
			validation.Required.Error("Email is required"),
			is.Email.Error("Email invalidate"),
		),

		validation.Field(
			&payload.FullName,
			validation.Required.Error("Fullname is required"),
			validation.Length(5, 30).Error("Fullname is length: 5 -> 30"),
		),

		validation.Field(
			&payload.Gender,
			validation.Required.Error("Gender is required"),
			validation.In("female", "male"),
		),

		validation.Field(
			&payload.DateOfBirth,
			validation.Required.Error("DateOfBirth is required"),
		),
	)
}

func (payload UserLogin) Validate() error {
	return validation.ValidateStruct(&payload,

		validation.Field(
			&payload.Username,
			validation.Required.Error("Username is required"),
			validation.Length(5, 30).Error("UserName is length: 5 -> 30"),
		),

		validation.Field(
			&payload.Password,
			validation.Required.Error("Password is required"),
			validation.Length(5, 30).Error("Password is length: 5 -> 30"),
		),
	)
}

func (body UserChangePassword) Validate() error {
	return validation.ValidateStruct(&body,

		validation.Field(
			&body.CurrentPassword,
			validation.Required.Error("CurrentPassword is required"),
			validation.Length(5, 30).Error("CurrentPassword is length: 5 -> 30"),
		),

		validation.Field(
			&body.NewPassword,
			validation.Required.Error("NewPassword is required"),
			validation.Length(5, 30).Error("NewPassword is length: 5 -> 30"),
		),
	)
}

func (body UserUpdate) Validate() error {
	return validation.ValidateStruct(&body,
		validation.Field(
			&body.FullName,
			validation.Required.Error("FullName is required"),
			validation.Length(5, 30).Error("FullName is length: 5 -> 30"),
		),

		validation.Field(
			&body.Email,
			validation.Required.Error("Email is required"),
			validation.Length(5, 30).Error("Email is length: 5 -> 30"),
		),
		validation.Field(
			&body.Phone,
			validation.Required.Error("Phone is required"),
			validation.Length(5, 30).Error("Phone is length: 5 -> 30"),
		),
		validation.Field(
			&body.DateOfBirth,
			validation.Required.Error("DateOfBirth is required"),
		),
		validation.Field(
			&body.Address,
			validation.Required.Error("Address is required"),
			validation.Length(5, 30).Error("Address is length: 5 -> 30"),
		),
	)
}
