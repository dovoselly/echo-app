package util

import "context"

var (
	CreateSuccessFully            = "Create successfully"
	YouCannotEditThis             = "You cannot edit this"
	UpdateSuccessFully            = "Update successfully"
	UserNotFound                  = "User not found"
	DeleteSuccessFully            = "Delete successfully"
	InvalidData                   = "Invalid data"
	EmailOrPasswordWrong          = "Email or password wrong"
	EmailOrUsernameIsAlreadyExist = "Email or username already exists"
	LoginSuccessFully             = "Login successfully"
	InvalidToken                  = "Invalid token"

	Ctx = context.Background()

	// Auth
	NOT_EXIST_USER        = "user not existed in database"
	WRONG_PASSWORD        = "Wrong password"
	GENERATE_TOKEN_FAILED = "Generate token failed"

	// Status User
	USER_STATUS_ACTIVE  = "ACTIVE"
	USER_STATUS_BLOCKED = "BLOCKED"

	//
	CURRENT_PASSWORD_INCORRECT = "CurrentPassword is incorrect"

	// Order
	ORDER_STATUS_PENDING    = "PENDING"
	ORDER_STATUS_CONFIRMED  = "CONFIRMED"
	ORDER_STATUS_DELIVERING = "DELIVERING"
	ORDER_STATUS_DELIVERED  = "DELIVERED"
	ORDER_STATUS_CANCEL     = "CANCEL"
)
