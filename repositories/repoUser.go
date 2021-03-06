package repositories

import "assyifa/models"

type UserRepositories interface {
	GetData() ([]*models.Users, error)
	AddUser(users models.Users) error
}
