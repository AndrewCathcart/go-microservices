package services

import (
	"net/http"

	"github.com/andrewcathcart/go-microservices/mvc/domain"
	"github.com/andrewcathcart/go-microservices/mvc/utils"
)

type itemsService struct {
}

var (
	ItemsService itemsService
)

// GetItem returns an item
func (s *itemsService) GetItem(itemID string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "Implement me",
		StatusCode: http.StatusInternalServerError,
		Code:       "not_implemented",
	}
}
