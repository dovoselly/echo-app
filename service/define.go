package service

import "echo-app/dao"

var (
	userDAO    = dao.User{}
	productDAO = dao.Product{}
	reviewDAO  = dao.Review{}
	replyDAO   = dao.Reply{}
)
