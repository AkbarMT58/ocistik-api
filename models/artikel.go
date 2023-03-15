package models

import "database/sql"

type Artikel struct {
	Id               int            `json:"id"`
	Title            string         `json:"title"`
	Description      string         `json:"description"`
	Konten           string         `json:"konten"`
	Slug             string         `json:"slug"`
	Slug_id          int            `json:"slug_id"`
	Headline_id      int            `json:"headline_id"`
	Picture          string         `json:"picture"`
	Penulis          string         `json:"penulis"`
	Tag_name         string         `json:"tag_name"`
	Meta_title       string         `json:"meta_title"`
	Meta_description string         `json:"meta_description"`
	Meta_keyword     string         `json:"meta_keyword"`
	Created_date     sql.NullString `json:"created_date"`
	Updated_date     sql.NullString `json:"updated_date"`
	Deleted_date     sql.NullString `json:"Deleted_date"`
}

type Responses struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Artikel
}

type Thelast_id struct {
	Last_id int `json:"last_id"`
}
