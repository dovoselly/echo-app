package controller

import (
	"echo-app/model"
	"echo-app/utils"

	"github.com/labstack/echo/v4"
)

type Reply struct{}

func (Reply) CreateReply(c echo.Context) error {
	reviewId := c.Param("id")

	payloadInterface := c.Get("payload")
	payload := payloadInterface.(model.CreateReply)

	userId, err := utils.GetUserId(c)
	if err != nil {
		return utils.Response400(c, nil, utils.InvalidData)
	}

	InsertedId, err := replyService.CreateReply(userId, reviewId, payload)
	if err != nil {
		return utils.Response400(c, nil, utils.InvalidData)
	}
	return utils.Response200(c, InsertedId, utils.UpdateSuccessFully)
}

func (Reply) UpdateReply(c echo.Context) error {
	replyId := c.Param("id")

	userId, err := utils.GetUserId(c)
	if err != nil {
		return utils.Response400(c, nil, utils.InvalidData)
	}

	body := c.Get("body").(model.CreateReply)

	results, err := replyService.UpdateReply(userId, replyId, body)
	if err != nil || results.MatchedCount == 0 {
		return utils.Response400(c, nil, utils.InvalidData)
	}
	return utils.Response200(c, replyId, utils.UpdateSuccessFully)
}

func (Reply) DeleteReply(c echo.Context) error {
	replyId := c.Param("id")

	userId, err := utils.GetUserId(c)
	if err != nil {
		return utils.Response400(c, nil, utils.InvalidData)
	}

	results, err := replyService.DeleteReply(userId, replyId)
	if err != nil || results.DeletedCount == 0 {
		return utils.Response400(c, nil, utils.InvalidData)
	}
	return utils.Response200(c, nil, utils.DeleteSuccessFully)
}
