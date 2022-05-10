package services

//var ctx = context.TODO()
//
//func Register(insertData models.User) string {
//	_, err := database.GetUserCol().InsertOne(ctx, insertData)
//	if err != nil {
//		return err.Error()
//	}
//	return utils.CreateSuccessFully
//}
//
//func GetUserByEmail(email string, username string) models.User {
//	var user models.User
//	err := database.GetUserCol().FindOne(ctx, bson.M{"$or": []bson.M{{"email": email}, {username: username}}}).Decode(&user)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	return user
//}

//
//func AllUsers(page int, limit int) []models.User {
//	var allUsers []models.User
//
//	// options query
//	optionsQuery := new(options.FindOptions)
//	optionsQuery.SetSkip(int64(page * limit))
//	optionsQuery.SetLimit(int64(limit))
//
//	// Find users
//	cursor, err := database.GetUserCol().Find(ctx, bson.M{}, optionsQuery)
//	if err != nil {
//		fmt.Println(err.Error())
//		return allUsers
//	}
//
//	// Decode found documents
//	if err := cursor.All(ctx, &allUsers); err != nil {
//		fmt.Println(err.Error())
//		return allUsers
//	}
//
//	return allUsers
//}
//
//func GetUserById(id primitive.ObjectID) models.User {
//	var user models.User
//	err := database.GetUserCol().FindOne(ctx, bson.M{"_id": id}).Decode(&user)
//	if err != nil {
//		fmt.Println(err.Error())
//		return user
//	}
//	return user
//}
//
//func UpdateUserById(id primitive.ObjectID, insertData models.UserUpdate) string {
//	fmt.Println(id, insertData)
//	result, err := database.GetUserCol().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": insertData})
//	fmt.Println(*result)
//	if err != nil {
//		return err.Error()
//	}
//	if result.MatchedCount == 0 {
//		return utils.UserNotFound
//	}
//	return utils.UpdateSuccessFully
//}
//
//func DeleteUserById(id primitive.ObjectID) string {
//	result, err := database.GetUserCol().DeleteOne(ctx, bson.M{"_id": id})
//	if err != nil {
//		return err.Error()
//	}
//	if result.DeletedCount == 0 {
//		return utils.UserNotFound
//	}
//	return utils.DeleteSuccessFully
//}
