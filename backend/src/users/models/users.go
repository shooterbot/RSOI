package models

type UserAuthData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username string `json:"username"`
	UUID     string `json:"UUID"`
	JWT      string `json:"JWT"`
}
