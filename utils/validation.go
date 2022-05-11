package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

// Convert ID string ->  primitive.ObjectID
func ValidateObjectID(id string) error {
	// ObjectIDFromHex
	_, err := primitive.ObjectIDFromHex(id)

	// if err
	if err != nil {
		return err
	}

	return nil
}
