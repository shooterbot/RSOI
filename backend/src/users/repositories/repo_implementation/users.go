package repo_implementation

import (
	"RSOI/src/database/pgdb"
	"RSOI/src/users/models"
	"fmt"
)

const (
	createUser = `insert into users(username, password) values($1, crypt($2, gen_salt('bf')));`
	loginUser  = `select from users where username=$1 and password=crypt($2, password);`
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

	return data != nil, err
}
