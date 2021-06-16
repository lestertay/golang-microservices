package services

import "github.com/tahkiu/golang-microservices/mvc/domain"

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}
