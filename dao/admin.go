package dao

import (
	"context"
	"echo-app/database"
	"echo-app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct{}

// FindByUsername
func (Admin) FindByUsername(username string) (model.Admin, error) {
	var (
		adminCol = database.AdminCol()
		ctx      = context.Background()
		admin    model.Admin
	)

	// find
	filter := bson.M{
		"username": username,
	}
	if err := adminCol.FindOne(ctx, filter).Decode(&admin); err != nil {
		return admin, err
	}

	//if err != nil {
	//	return admin, err
	//}

	return admin, nil
}

func (Admin) InitAdminAccount() {
	var (
		adminCol = database.AdminCol()
		ctx      = context.Background()
	)

	count, _ := adminCol.CountDocuments(ctx, bson.D{})

	if count == 0 {
		admin := model.Admin{
			ID:             primitive.NewObjectID(),
			Email:          "admin123@gmail.com",
			Username:       "admin123",
			HashedPassword: "123456",
			FullName:       "admin",
			DateOfBirth:    "2022-05-10",
			Gender:         "Male",
			Phone:          "0779925153",
		}
		adminCol.InsertOne(ctx, admin)
	}
}

func (Admin) ProfileFindByID(ID string) (model.Admin, error) {
	var (
		adminCol     = database.AdminCol()
		ctx          = context.Background()
		adminProfile model.Admin
	)

	// objectID
	objID, _ := primitive.ObjectIDFromHex(ID)

	// find profile
	filter := bson.M{"_id": objID}
	if err := adminCol.FindOne(ctx, filter).Decode(&adminProfile); err != nil {
		return adminProfile, err
	}

	//// if err
	//if err != nil {
	//	return adminProfile, err
	//}

	return adminProfile, nil

}

func (Admin) UpdateMyProfile(ID string, newProfile model.Admin) error {
	var (
		adminCol = database.AdminCol()
		ctx      = context.Background()
	)

	objID, _ := primitive.ObjectIDFromHex(ID)
	update := model.Admin{
		FullName:    newProfile.FullName,
		DateOfBirth: newProfile.DateOfBirth,
		Gender:      newProfile.Gender,
		Phone:       newProfile.Phone,
	}

	// UpdateOne
	if _, err := adminCol.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": update}); err != nil {
		return err
	}

	//if err != nil {
	//	return err
	//}

	return nil
}
