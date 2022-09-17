package usecases

import "RSOI/src/users/models"

type IUsersUsecase interface {
	CreateUser(user *models.User) error
	LoginUser(user *models.User) (bool, error)
}