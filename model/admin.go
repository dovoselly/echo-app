package model

type (
	Admin struct {
		ID             string ` json:"_id" bson:"_id"`
		Email          string ` json:"email" bson:"email"`
		Username       string ` json:"username" bson:"username"`
		HashedPassword string ` json:"hashedPassword" bson:"hashedPassword"`
		FullName       string ` json:"fullName" bson:"fullName"`
		DateOfBirth    string ` json:"dateOfBirth" bson:"dateOfBirth"`
		Avatar         string `json:"avatar" bson:"avatar"`
		Gender         string ` json:"gender" bson:"gender"`
		Phone          string ` json:"phone" bson:"phone"`
		CreatedAt      string ` json:"createdAt" bson:"createdAt"`
		UpdatedAt      string ` json:"updatedAt" bson:"updatedAt"`
	}
)
