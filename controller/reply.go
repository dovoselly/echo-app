package controller

import (
	"echo-app/models"
	"echo-app/service"
	"echo-app/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateReply(c echo.Context) error {
	reviewIdString := c.Param("id")
	reviewId, err := primitive.ObjectIDFromHex(reviewIdString)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	payloadInterface := c.Get("payload")
	payload := payloadInterface.(models.CreateReply)

	userId, err := util.GetUserId(c)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	err = service.CreateReply(userId, reviewId, payload)
	return util.Response200(c, nil, err.Error())
}

func UpdateReply(c echo.Context) error {
	replyIdString := c.Param("id")
	replyId, err := primitive.ObjectIDFromHex(replyIdString)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	userId, err := util.GetUserId(c)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	bodyInterface := c.Get("body")
	body, ok := bodyInterface.(models.CreateReply)
	if !ok {
		return util.Response400(c, nil, util.InvalidData)
	}

	results, err := service.UpdateReply(userId, replyId, body)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}
	if results.MatchedCount == 0 {
		return util.Response400(c, nil, util.InvalidData)
	}

	return util.Response200(c, nil, util.UpdateSuccessFully)
}

func DeleteReply(c echo.Context) error {
	replyIdString := c.Param("id")
	replyId, err := primitive.ObjectIDFromHex(replyIdString)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	userId, err := util.GetUserId(c)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	results, err := service.DeleteReply(userId, replyId)

	if results.DeletedCount == 0 {
		return util.Response400(c, nil, util.InvalidData)
	}

	return util.Response200(c, nil, util.DeleteSuccessFully)
}
