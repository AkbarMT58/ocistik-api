package models

import "database/sql"

type Artikel struct {
	Id           int            `json:"id"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	Slug         string         `json:"slug"`
	Headline_id  int            `json:"headline_id"`
	Picture      string         `json:"picture"`
	Created_date sql.NullString `json:"created_date"`
	Updated_date sql.NullString `json:"updated_date"`
	Deleted_date sql.NullString `json:"Deleted_date"`
}

type Responses struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Artikel
}
