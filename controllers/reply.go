package controllers

import (
	"echo-app/models"
	"echo-app/services"
	"echo-app/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateReply(c echo.Context) error {
	reviewIdString := c.Param("id")
	reviewId, err := primitive.ObjectIDFromHex(reviewIdString)
	if err != nil {
		return utils.Response400(c, nil, utils.InvalidData)
	}

	payloadInterface := c.Get("payload")
	payload := payloadInterface.(models.CreateReply)

	userId, err := utils.GetUserId(c)
	if err != nil {
		return utils.Response400(c, nil, utils.InvalidData)
	}

	err = services.CreateReply(userId, reviewId, payload)
	return utils.Response200(c, nil, err.Error())
}

func UpdateReply(c echo.Context) error {
	replyIdString := c.Param("id")
	replyId, err := primitive.ObjectIDFromHex(replyIdString)
	if err != nil {
		return utils.Response400(c, nil, utils.InvalidData)
	}

	userId, err := utils.GetUserId(c)
	if err != nil {
		return utils.Response400(c, nil, utils.InvalidData)
	}

	bodyInterface := c.Get("body")
	body, ok := bodyInterface.(models.CreateReply)
	if !ok {
		return utils.Response400(c, nil, utils.InvalidData)
	}

	results, err := services.UpdateReply(userId, replyId, body)
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}
	if results.MatchedCount == 0 {
		return utils.Response400(c, nil, utils.InvalidData)
	}

	return utils.Response200(c, nil, utils.UpdateSuccessFully)
}

func DeleteReply(c echo.Context) error {
	replyIdString := c.Param("id")
	replyId, err := primitive.ObjectIDFromHex(replyIdString)
	if err != nil {
		return utils.Response400(c, nil, utils.InvalidData)
	}

	userId, err := utils.GetUserId(c)
	if err != nil {
		return utils.Response400(c, nil, utils.InvalidData)
	}

	results, err := services.DeleteReply(userId, replyId)

	if results.DeletedCount == 0 {
		return utils.Response400(c, nil, utils.InvalidData)
	}

	return utils.Response200(c, nil, utils.DeleteSuccessFully)
}
