package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/tahkiu/golang-microservices/mvc/services"
	"github.com/tahkiu/golang-microservices/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		apiErr := utils.ApiError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		error_json, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(error_json)
		return
	}

	user, apiErr := services.GetUser(userId)
	error_json, _ := json.Marshal(apiErr)
	if err != nil {
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(error_json)
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
