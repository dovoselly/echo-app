package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
)

type (
	User struct {
		ID        primitive.ObjectID `param:"id" json:"_id" bson:"_id"`
		Name      string             `param:"name" query:"name" json:"name" bson:"name"`
		Dob       string             `param:"dob" query:"dob" json:"dob" bson:"dob"`
		Email     string             `param:"email" query:"email" json:"email" bson:"email"`
		Password  string             `json:"password" bson:"password"`
		CreatedAt string             `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	}

	UserLogin struct {
		Email    string `param:"email" query:"email" json:"email" bson:"email"`
		Password string `json:"password" bson:"password"`
	}

	UserUpdate struct {
		Name string `param:"name" query:"name" json:"name" bson:"name"`
		Dob  string `param:"dob" query:"dob" json:"dob" bson:"dob"`
	}
)

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Length(2, 50)),
		validation.Field(&u.Dob, validation.Required),
		validation.Field(&u.Email, validation.Length(8, 50)),
		validation.Field(&u.Password, validation.Match(regexp.MustCompile("^[a-zA-Z0-9_!@#$%^&*()~+=,./?;:]{8,50}$"))))
}

func (u UserUpdate) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Length(2, 50)),
		validation.Field(&u.Dob, validation.Required),
	)
}

func (u UserLogin) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Length(8, 50)),
		validation.Field(&u.Password, validation.Match(regexp.MustCompile("^[a-zA-Z0-9_!@#$%^&*()`~+=,./:;]{8,50}$"))),
	)
}
