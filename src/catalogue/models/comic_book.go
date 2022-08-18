package models

type Book struct {
	Id        int    `json:"id"`
	Uuid      string `json:"Uuid"`
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
}
