package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/andrewcathcart/go-microservices/mvc/services"
	"github.com/andrewcathcart/go-microservices/mvc/utils"
)

// GetUser returns a User with the userID matching the user_id query param from the URL
func GetUser(res http.ResponseWriter, req *http.Request) {
	// get userID from query params
	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userID)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
