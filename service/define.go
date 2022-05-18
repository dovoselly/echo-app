package service

import "echo-app/dao"

var (
	adminDao    = dao.Admin{}
	categoryDao = dao.Category{}
	brandDao    = dao.Brand{}
	userDAO    = dao.User{}
	productDAO = dao.Product{}
	reviewDAO  = dao.Review{}
)
