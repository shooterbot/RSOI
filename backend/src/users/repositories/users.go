package repositories

import "RSOI/src/users/models"

type IUsersRepository interface {
	CreateUser(user *models.User) error
	LoginUser(user *models.User) (bool, error)
}
