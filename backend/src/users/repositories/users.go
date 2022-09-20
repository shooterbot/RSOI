package repositories

import "RSOI/src/users/models"

type IUsersRepository interface {
	CreateUser(user *models.User) error
	LoginUser(user *models.User) (bool, error)
	GetUserPreferences(uuid string) (models.PreferencesList, error)
	GetUserUuid(username string) (string, error)
	GetUserPreference(userUuid string, bookUuid string) (string, error)
	SetLike(userUuid string, bookUuid string) error
	RemoveDislike(userUuid string, bookUuid string) error
	SetDislike(userUuid string, bookUuid string) error
	RemoveLike(userUuid string, bookUuid string) error
}
