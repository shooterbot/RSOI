package uc_implementation

import (
	"RSOI/src/users/models"
	"RSOI/src/users/repositories"
	"errors"
)

type UsersUsecase struct {
	ur repositories.IUsersRepository
}

func NewUsersUsecase(repo repositories.IUsersRepository) *UsersUsecase {
	return &UsersUsecase{ur: repo}
}

func (uc *UsersUsecase) CreateUser(user *models.UserAuthData) error {
	return uc.ur.CreateUser(user)
}

func (uc *UsersUsecase) LoginUser(user *models.UserAuthData) (*models.User, error) {
	var res *models.User = nil
	uuid, err := uc.ur.LoginUser(user)
	if err == nil {
		res = &models.User{
			Username: user.Username,
			UUID:     uuid,
		}
	}
	return res, err
}

func (uc *UsersUsecase) GetUserPreferences(uuid string) (models.PreferencesList, error) {
	return uc.ur.GetUserPreferences(uuid)
}

func (uc *UsersUsecase) SetUserScore(username string, bookUuid string, score string) (bool, error) {
	var pref string
	var err error
	if score != "like" && score != "dislike" {
		err = errors.New("Bad score given")
		return false, err
	}

	userUuid, err := uc.ur.GetUserUuid(username)
	if err != nil {
		err = errors.New("Bad username given")
		return false, err
	}

	pref, err = uc.ur.GetUserPreference(userUuid, bookUuid)
	if err != nil {
		err = errors.New("Failed to get user preference")
		return false, err
	}
	if pref == score {
		err = errors.New("Given score equals the old score")
		return false, err
	}
	changed := pref != "none"
	if score == "like" {
		err = uc.ur.SetLike(userUuid, bookUuid)
		if err == nil && changed {
			err = uc.ur.RemoveDislike(userUuid, bookUuid)
		}
	} else {
		err = uc.ur.SetDislike(userUuid, bookUuid)
		if err == nil && changed {
			err = uc.ur.RemoveLike(userUuid, bookUuid)
		}
	}
	return changed, err
}
