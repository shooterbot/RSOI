package models

type PreferencesList struct {
	Likes    []int
	Dislikes []int
}

type Book struct {
	Id        int
	Uuid      string
	Name      string
	Publisher string
	Year      int
	Rating    int
	Status    bool
	Tags      []string
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
