package services

import (
	"github.com/andrewcathcart/go-microservices/mvc/models"
	"github.com/andrewcathcart/go-microservices/mvc/utils"
)

func GetUser(userID int64) (*models.User, *utils.ApplicationError) {
	return models.GetUser(userID)
}
