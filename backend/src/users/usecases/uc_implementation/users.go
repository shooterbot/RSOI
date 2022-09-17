package uc_implementation

import (
	"RSOI/src/users/models"
	"RSOI/src/users/repositories"
)

type UsersUsecase struct {
	ur repositories.IUsersRepository
}

func (uc *UsersUsecase) CreateUser(user *models.User) error {
	return uc.ur.CreateUser(user)
}

func (uc *UsersUsecase) LoginUser(user *models.User) (bool, error) {
	return uc.ur.LoginUser(user)
}
