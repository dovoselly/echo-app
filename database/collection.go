package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	admins     = "admins"
	users      = "users"
	categories = "categories"
	brands     = "brands"
	products   = "products"
	reviews    = "reviews"
	replies    = "replies"
	carts      = "carts"
	cartItems  = "cartItems"
	orders     = "orders"
	orderItems = "orderItems"
	// ...
)

func UserCol() *mongo.Collection {
	return db.Collection(users)
}

func AdminCol() *mongo.Collection {
	return db.Collection(admins)
}

func CategoryCol() *mongo.Collection {
	return db.Collection(categories)
}

// ...
