package usecases

import "RSOI/src/users/models"

type IUsersUsecase interface {
	CreateUser(user *models.UserAuthData) error
	LoginUser(user *models.UserAuthData) (*models.User, error)
	GetUserPreferences(uuid string) (models.PreferencesList, error)
	SetUserScore(username string, uuid string, score string) (bool, error)
}
