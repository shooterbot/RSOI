package repo_implementation

import (
	"RSOI/src/database/pgdb"
	"RSOI/src/users/models"
	"RSOI/src/utility"
	"errors"
	"fmt"
)

const (
	createUser       = `insert into users(username, password) values($1, crypt($2, gen_salt('bf')));`
	loginUser        = `select uuid from users where username=$1 and password=crypt($2, password);`
	getUserLikes     = `select liked from users_likes where user_id in (select id from users where uid=$1);`
	getUserDislikes  = `select disliked from users_dislikes where user_id in (select id from users where uid=$1);`
	getUserUuid      = `select uid from users where username = $1`
	checkUserLike    = `select from users_likes where user_id in (select id from users where uid=$1) and liked=$2`
	checkUserDislike = `select from users_dislikes where user_id in (select id from users where uid=$1) and disliked=$2`
	setLike          = `insert into users_likes(user_id, liked) values((select first_value(id) over (order by id) from users where uid = $1), $2);`
	setDislike       = `insert into users_dislikes(user_id, disliked) values((select first_value(id) over (order by id) from users where uid = $1), $2);`
	removeLike       = `delete from users_likes where user_id in (select id from users where uid=$1) and liked = $2`
	removeDislike    = `delete from users_dislikes where user_id in (select id from users where uid=$1) and disliked = $2`
)

type UsersRepository struct {
	db *pgdb.DBManager
}

func NewUsersRepository(manager *pgdb.DBManager) *UsersRepository {
	return &UsersRepository{db: manager}
}

func (ur *UsersRepository) CreateUser(user *models.UserAuthData) error {
	_, err := ur.db.Exec(createUser, user.Username, user.Password)
	if err != nil {
		fmt.Printf("Failed to add new user to db\n")
	}
	return err
}

func (ur *UsersRepository) LoginUser(user *models.UserAuthData) (string, error) {
	data, err := ur.db.Query(loginUser, user.Username, user.Password)
	if err != nil {
		fmt.Printf("Failed to login user via db\n")
	}

	res := ""
	if len(data) == 1 {
		res = utility.BytesToUid(data[0][0])
	}

	return res, err
}

func (ur *UsersRepository) GetUserPreferences(uuid string) (models.PreferencesList, error) {
	res := models.PreferencesList{
		Likes:    make([]int, 0),
		Dislikes: make([]int, 0),
	}
	data, err := ur.db.Query(getUserLikes, uuid)
	if err != nil {
		fmt.Printf("Failed to get likes from db\n")
	}
	for _, row := range data {
		res.Likes = append(res.Likes, utility.BytesToInt(row[0]))
	}
	data, err = ur.db.Query(getUserDislikes, uuid)
	if err != nil {
		fmt.Printf("Failed to get dislikes from db\n")
	}
	for _, row := range data {
		res.Dislikes = append(res.Dislikes, utility.BytesToInt(row[0]))
	}
	return res, err
}

func (ur *UsersRepository) GetUserUuid(username string) (string, error) {
	data, err := ur.db.Query(getUserUuid, username)
	if err != nil {
		fmt.Printf("Failed to get user uuid from db\n")
		return "", err
	}
	return utility.BytesToUid(data[0][0]), nil
}

func (ur *UsersRepository) GetUserPreference(userUuid string, bookUuid string) (string, error) {
	data, err := ur.db.Query(checkUserLike, userUuid, bookUuid)
	if err != nil {
		return "none", err
	}
	if len(data) != 0 {
		return "like", nil
	} else {
		data, err = ur.db.Query(checkUserDislike, userUuid, bookUuid)
		if err != nil {
			return "none", err
		}
		if len(data) != 0 {
			return "dislike", nil
		} else {
			return "none", nil
		}
	}
}

func (ur *UsersRepository) SetLike(userUuid string, bookUuid string) error {
	affected, err := ur.db.Exec(setLike, userUuid, bookUuid)
	if err == nil && affected == 0 {
		err = errors.New("UserAuthData's like is already set")
	}
	return err
}

func (ur *UsersRepository) SetDislike(userUuid string, bookUuid string) error {
	affected, err := ur.db.Exec(setDislike, userUuid, bookUuid)
	if err == nil && affected == 0 {
		err = errors.New("UserAuthData's dislike is already set")
	}
	return err
}

func (ur *UsersRepository) RemoveLike(userUuid string, bookUuid string) error {
	affected, err := ur.db.Exec(removeLike, userUuid, bookUuid)
	if err == nil && affected == 0 {
		err = errors.New("UserAuthData's like was not set")
	}
	return err
}

func (ur *UsersRepository) RemoveDislike(userUuid string, bookUuid string) error {
	affected, err := ur.db.Exec(removeDislike, userUuid, bookUuid)
	if err == nil && affected == 0 {
		err = errors.New("UserAuthData's dislike was not set")
	}
	return err
}
