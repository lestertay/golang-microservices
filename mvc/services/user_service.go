package services

import (
	"github.com/tahkiu/golang-microservices/mvc/domain"
	"github.com/tahkiu/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApiError) {
	return domain.GetUser(userId)
}
