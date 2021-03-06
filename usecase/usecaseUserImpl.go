package usecase

import (
	"assyifa/models"
	"assyifa/repositories"
	"log"
)

type UserUsecaseImpl struct {
	userRepo repositories.UserRepositories
}

func (User UserUsecaseImpl) GetAllData() ([]*models.Users, error) {
	data, err := User.userRepo.GetData()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (User UserUsecaseImpl) AddUser(Users models.Users) error {
	err := User.userRepo.AddUser(Users)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func InitUsecaseImpl(userRepo repositories.UserRepositories) UserUsecase {
	return &UserUsecaseImpl{userRepo}
}
