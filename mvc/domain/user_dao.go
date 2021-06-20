package domain

import (
	"fmt"
	"net/http"

	"github.com/tahkiu/golang-microservices/mvc/utils"
)

var users = map[int64]*User{
	123: {
		Id:        123,
		FirstName: "tah",
		LastName:  "kiu",
		Email:     "kiukiu@gmail.com",
	},
}

func GetUser(userId int64) (*User, *utils.ApiError) {
	user := users[userId]
	if user == nil {
		return nil, &utils.ApiError{
			Message:    fmt.Sprintf("user %d not found", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not found",
		}
	}

	return user, nil
}
