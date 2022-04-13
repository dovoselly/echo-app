package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID   primitive.ObjectID `param:"id,omitempty" query:"id,omitempty" json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `param:"name,omitempty" query:"name,omitempty" json:"name,omitempty" bson:"name,omitempty"`
	Age  uint16             `param:"age,omitempty" query:"age,omitempty" json:"age,omitempty" bson:"age,omitempty"`
}
