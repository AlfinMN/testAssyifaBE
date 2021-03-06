package usecase

import "assyifa/models"

type UserUsecase interface {
	GetAllData() ([]*models.Users, error)
	AddUser(users models.Users) error
}
