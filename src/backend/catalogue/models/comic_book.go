package models

type Book struct {
	Id        int      `json:"id"`
	Uuid      string   `json:"Uuid"`
	Name      string   `json:"name"`
	Publisher string   `json:"publisher"`
	Year      int      `json:"year"`
	Rating    int      `json:"rating"`
	Status    bool     `json:"status"`
	Tags      []string `json:"tags"`
}
