package models

type PreferencesList struct {
	Likes    []int `json:"likes"`
	Dislikes []int `json:"dislikes"`
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

type RecomendationsInfo struct {
	Books []Book          `json:"books"`
	Prefs PreferencesList `json:"prefs"`
}
