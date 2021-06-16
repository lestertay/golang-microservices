package domain

import (
	"fmt"
)

var users = map[int64]*User{
	123: {
		Id:        123,
		FirstName: "tah",
		LastName:  "kiu",
		Email:     "kiukiu@gmail.com",
	},
}

func GetUser(userId int64) (*User, error) {
	user := users[userId]
	if user == nil {
		return nil, fmt.Errorf("user %v was not found", userId)
	}

	return user, nil
}
