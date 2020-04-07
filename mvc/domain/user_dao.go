package domain

import (
	"fmt"
	"net/http"

	"github.com/andrewcathcart/go-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Andy", LastName: "Cathcart", Email: "test@gmail.com"},
	}

	UserDao userDaoInterface
)

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

func init() {
	UserDao = &userDao{}
}

type userDao struct{}

func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v does not exist", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
