package services

import (
	"github.com/andrewcathcart/go-microservices/mvc/domain"
	"github.com/andrewcathcart/go-microservices/mvc/utils"
)

type usersService struct {
}

var (
	UsersService usersService
)

func (u *usersService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userID)
}
