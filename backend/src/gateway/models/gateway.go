package models

type PreferencesList struct {
	Likes    []string `json:"likes"`
	Dislikes []string `json:"dislikes"`
}

type Book struct {
	Id        int      `json:"id"`
	Uuid      string   `json:"uuid"`
	Name      string   `json:"name"`
	Publisher string   `json:"publisher"`
	Year      int      `json:"year"`
	Rating    int      `json:"rating"`
	Status    bool     `json:"status"`
	Tags      []string `json:"tags"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RecomendationsInfo struct {
	Books []Book          `json:"books"`
	Prefs PreferencesList `json:"prefs"`
}

type Session struct {
	Username string `json:"username"`
	UUID     string `json:"UUID"`
	JWT      string `json:"JWT"`
}
