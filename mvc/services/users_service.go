package services

import (
	"github.com/andrewcathcart/go-microservices/mvc/domain"
	"github.com/andrewcathcart/go-microservices/mvc/utils"
)

func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
