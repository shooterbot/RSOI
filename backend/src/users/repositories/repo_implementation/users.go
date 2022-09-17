package repo_implementation

import (
	"RSOI/src/database/pgdb"
	"RSOI/src/users/models"
	"RSOI/src/utility"
	"fmt"
)

const (
	createUser      = `insert into users(username, password) values($1, crypt($2, gen_salt('bf')));`
	loginUser       = `select from users where username=$1 and password=crypt($2, password);`
	getUserLikes    = `select liked from users_likes where user_id in (select id from users where uid=$1);`
	getUserDislikes = `select disliked from users_dislikes where user_id in (select id from users where uid=$1);`
)

type UsersRepository struct {
	db *pgdb.DBManager
}

func NewUsersRepository(manager *pgdb.DBManager) *UsersRepository {
	return &UsersRepository{db: manager}
}

func (ur *UsersRepository) CreateUser(user *models.User) error {
	_, err := ur.db.Exec(createUser, user.Username, user.Password)
	if err != nil {
		fmt.Printf("Failed to add new user to db\n")
	}
	return err
}

func (ur *UsersRepository) LoginUser(user *models.User) (bool, error) {
	data, err := ur.db.Query(loginUser, user.Username, user.Password)
	if err != nil {
		fmt.Printf("Failed to login user via db\n")
	}

	return len(data) == 1, err
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
